package sdfs

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"k8s.io/klog/v2"
)

// ControllerServer controller server setting
type ControllerServer struct {
	Driver *Driver
	// Working directory for the provisioner to temporarily mount sdfs shares at
	workingMountDir string
}

// sdfsVolume is an internal representation of a volume
// created by the provisioner.
type sdfsVolume struct {
	// Volume id
	id string
	// Address of the SDFS server.
	// Matches paramServer.
	server string
	// Subdirectory of the SDFS server to create volumes under
	subDir string
	// size of volume
	size int64
	// client side dedupe is set
	dedupe bool
	// user name for login
	user string
	// password
	password string
	// shared folder
	sharedFolder bool
}

// Ordering of elements in the CSI volume id.
// ID is of the form {server}/{baseDir}/{subDir}.
// TODO: This volume id format limits baseDir and
// subDir to only be one directory deep.
// Adding a new element should always go at the end
// before totalIDElements
const (
	idServer = iota
	idBaseDir
	idSubDir
	totalIDElements // Always last
)

// CreateVolume create a volume
func (cs *ControllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	name := req.GetName()
	if len(name) == 0 {
		return nil, status.Error(codes.InvalidArgument, "CreateVolume name must be provided")
	}
	if err := cs.validateVolumeCapabilities(req.GetVolumeCapabilities()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	reqCapacity := req.GetCapacityRange().GetRequiredBytes()
	sdfsVol, err := cs.newSDFSVolume(name, reqCapacity, req.GetParameters())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	var volCap *csi.VolumeCapability
	if len(req.GetVolumeCapabilities()) > 0 {
		volCap = req.GetVolumeCapabilities()[0]
	}
	// Mount sdfs base share so we can create a subdirectory
	if err = cs.internalMount(ctx, sdfsVol, volCap); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to mount sdfs server: %v", err.Error())
	}
	defer func() {
		if err = cs.internalUnmount(ctx, sdfsVol); err != nil {
			klog.Warningf("failed to unmount sdfs server: %v", err.Error())
		}
	}()

	// Create subdirectory under base-dir
	// TODO: revisit permissions
	/*
		internalVolumePath := cs.getInternalVolumePath(sdfsVol)
		if err = os.Mkdir(internalVolumePath, 0777); err != nil && !os.IsExist(err) {
			return nil, status.Errorf(codes.Internal, "failed to make subdirectory: %v", err.Error())
		}
	*/
	return &csi.CreateVolumeResponse{Volume: cs.sdfsVolToCSI(sdfsVol)}, nil
}

// DeleteVolume delete a volume
func (cs *ControllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	volumeID := req.GetVolumeId()
	if volumeID == "" {
		return nil, status.Error(codes.InvalidArgument, "volume id is empty")
	}
	sdfsVol, err := cs.getSdfsVolFromID(volumeID)
	if err != nil {
		// An invalid ID should be treated as doesn't exist
		klog.Warningf("failed to get sdfs volume for volume id %v deletion: %v", volumeID, err)
		return &csi.DeleteVolumeResponse{}, nil
	}

	// Mount sdfs base share so we can delete the subdirectory
	if err = cs.internalMount(ctx, sdfsVol, nil); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to mount sdfs server: %v", err.Error())
	}
	defer func() {
		if err = cs.internalUnmount(ctx, sdfsVol); err != nil {
			klog.Warningf("failed to unmount sdfs server: %v", err.Error())
		}
	}()

	// Delete subdirectory under base-dir
	internalVolumePath := cs.getInternalVolumePath(sdfsVol)

	klog.V(2).Infof("Removing subdirectory at %v", internalVolumePath)
	if err = os.RemoveAll(internalVolumePath); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete subdirectory: %v", err.Error())
	}

	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *ControllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) ControllerGetVolume(ctx context.Context, req *csi.ControllerGetVolumeRequest) (*csi.ControllerGetVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID missing in request")
	}
	if req.GetVolumeCapabilities() == nil {
		return nil, status.Error(codes.InvalidArgument, "Volume capabilities missing in request")
	}

	// supports all AccessModes, no need to check capabilities here
	return &csi.ValidateVolumeCapabilitiesResponse{Message: ""}, nil
}

func (cs *ControllerServer) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

// ControllerGetCapabilities implements the default GRPC callout.
// Default supports all capabilities
func (cs *ControllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: cs.Driver.cscap,
	}, nil
}

func (cs *ControllerServer) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (cs *ControllerServer) validateVolumeCapabilities(caps []*csi.VolumeCapability) error {
	if len(caps) == 0 {
		return fmt.Errorf("volume capabilities must be provided")
	}

	for _, c := range caps {
		if err := cs.validateVolumeCapability(c); err != nil {
			return err
		}
	}
	return nil
}

func (cs *ControllerServer) validateVolumeCapability(c *csi.VolumeCapability) error {
	if c == nil {
		return fmt.Errorf("volume capability must be provided")
	}

	// Validate access mode
	accessMode := c.GetAccessMode()
	if accessMode == nil {
		return fmt.Errorf("volume capability access mode not set")
	}
	if !cs.Driver.cap[accessMode.Mode] {
		return fmt.Errorf("driver does not support access mode: %v", accessMode.Mode.String())
	}

	// Validate access type
	accessType := c.GetAccessType()
	if accessType == nil {
		return fmt.Errorf("volume capability access type not set")
	}
	return nil
}

// Mount sdfs server at base-dir
func (cs *ControllerServer) internalMount(ctx context.Context, vol *sdfsVolume, volCap *csi.VolumeCapability) error {
	targetPath := cs.getInternalMountPath(vol)

	if volCap == nil {
		volCap = &csi.VolumeCapability{
			AccessType: &csi.VolumeCapability_Mount{
				Mount: &csi.VolumeCapability_MountVolume{},
			},
		}
	}

	klog.V(4).Infof("internally mounting %v at %v", vol.server, targetPath)
	_, err := cs.Driver.ns.NodePublishVolume(ctx, &csi.NodePublishVolumeRequest{
		TargetPath: targetPath,
		VolumeContext: map[string]string{
			paramServer:   vol.server,
			paramUser:     vol.user,
			paramPassword: vol.password,
			paramDedupe:   strconv.FormatBool(vol.dedupe),
			paramShared:   strconv.FormatBool(vol.sharedFolder),
			"subdir":      vol.subDir,
		},
		VolumeCapability: volCap,
		VolumeId:         vol.id,
	})
	return err
}

// Unmount sdfs server at base-dir
func (cs *ControllerServer) internalUnmount(ctx context.Context, vol *sdfsVolume) error {
	targetPath := cs.getInternalMountPath(vol)

	// Unmount sdfs server at base-dir
	klog.V(4).Infof("internally unmounting %v", targetPath)
	_, err := cs.Driver.ns.NodeUnpublishVolume(ctx, &csi.NodeUnpublishVolumeRequest{
		VolumeId:   vol.id,
		TargetPath: cs.getInternalMountPath(vol),
	})
	return err
}

// Convert VolumeCreate parameters to an sdfsVolume
func (cs *ControllerServer) newSDFSVolume(name string, size int64, params map[string]string) (*sdfsVolume, error) {
	var (
		server   string
		user     string
		password string
		dedupe   bool
		shared   bool
	)

	// Validate parameters (case-insensitive).
	// TODO do more strict validation.
	for k, v := range params {
		switch strings.ToLower(k) {
		case paramServer:
			server = v
		case paramUser:
			user = v
		case paramPassword:
			password = v
		case paramDedupe:
			dedupe, _ = strconv.ParseBool(v)
		case paramShared:
			shared, _ = strconv.ParseBool(v)

		default:
			return nil, fmt.Errorf("invalid parameter %q", k)
		}
	}

	// Validate required parameters
	if server == "" {
		return nil, fmt.Errorf("%v is a required parameter", paramServer)
	}

	vol := &sdfsVolume{
		server:       server,
		subDir:       name,
		size:         size,
		user:         user,
		password:     password,
		dedupe:       dedupe,
		sharedFolder: shared,
	}
	vol.id = cs.getVolumeIDFromSdfsVol(vol)

	return vol, nil
}

// Get working directory for CreateVolume and DeleteVolume
func (cs *ControllerServer) getInternalMountPath(vol *sdfsVolume) string {
	// use default if empty
	if cs.workingMountDir == "" {
		cs.workingMountDir = "/tmp"
	}
	return filepath.Join(cs.workingMountDir, vol.subDir)
}

// Get internal path where the volume is created
// The reason why the internal path is "workingDir/subDir/subDir" is because:
//   * the semantic is actually "workingDir/volId/subDir" and volId == subDir.
//   * we need a mount directory per volId because you can have multiple
//     CreateVolume calls in parallel and they may use the same underlying share.
//     Instead of refcounting how many CreateVolume calls are using the same
//     share, it's simpler to just do a mount per request.
func (cs *ControllerServer) getInternalVolumePath(vol *sdfsVolume) string {
	return filepath.Join(cs.getInternalMountPath(vol), vol.subDir)
}

// Convert into sdfsVolume into a csi.Volume
func (cs *ControllerServer) sdfsVolToCSI(vol *sdfsVolume) *csi.Volume {
	return &csi.Volume{
		CapacityBytes: 0, // by setting it to zero, Provisioner will use PVC requested size as PV size
		VolumeId:      vol.id,
		VolumeContext: map[string]string{
			paramServer:   vol.server,
			paramDedupe:   strconv.FormatBool(vol.dedupe),
			paramUser:     vol.user,
			paramPassword: vol.password,
			paramShared:   strconv.FormatBool(vol.sharedFolder),
		},
	}
}

// Given a sdfsVolume, return a CSI volume id
func (cs *ControllerServer) getVolumeIDFromSdfsVol(vol *sdfsVolume) string {
	idElements := make([]string, totalIDElements)
	idElements[idServer] = strings.Trim(vol.server, "/")
	idElements[idBaseDir] = ""
	idElements[idSubDir] = strings.Trim(vol.subDir, "/")
	return strings.Join(idElements, "/")
}

// Given a CSI volume id, return a sdfsVolume
func (cs *ControllerServer) getSdfsVolFromID(id string) (*sdfsVolume, error) {
	volRegex := regexp.MustCompile("^([^/]+)/(.*)/([^/]+)$")
	tokens := volRegex.FindStringSubmatch(id)
	if tokens == nil {
		return nil, fmt.Errorf("could not split %q into server, and subDir", id)
	}

	return &sdfsVolume{
		id:     id,
		server: tokens[1],
		subDir: tokens[2],
	}, nil
}
