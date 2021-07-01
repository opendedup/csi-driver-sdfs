package sdfs

import (
	"fmt"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"k8s.io/klog/v2"
	mount "k8s.io/mount-utils"
)

type Driver struct {
	name    string
	nodeID  string
	version string

	endpoint string

	perm *uint32

	//ids *identityServer
	ns          *NodeServer
	cap         map[csi.VolumeCapability_AccessMode_Mode]bool
	cscap       []*csi.ControllerServiceCapability
	nscap       []*csi.NodeServiceCapability
	volumeLocks *VolumeLocks
}

const (
	DriverName = "sdfs.csi.k8s.io"
	// Address of the SDFS server
	paramServer   = "server"   //the server url
	paramDedupe   = "dedupe"   //Set clientside dedupe enabled
	paramUser     = "user"     //The username to autheticate to sdfs
	paramPassword = "password" //The password to authenticate to sdfs
)

var (
	version = "0.0.1"
)

func NewSDFSdriver(nodeID, endpoint string, perm *uint32) *Driver {
	klog.Infof("Driver: %v version: %v", DriverName, version)

	n := &Driver{
		name:     DriverName,
		version:  version,
		nodeID:   nodeID,
		endpoint: endpoint,
		cap:      map[csi.VolumeCapability_AccessMode_Mode]bool{},
		perm:     perm,
	}

	vcam := []csi.VolumeCapability_AccessMode_Mode{
		csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER,
		csi.VolumeCapability_AccessMode_SINGLE_NODE_READER_ONLY,
		csi.VolumeCapability_AccessMode_MULTI_NODE_READER_ONLY,
		csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER,
		csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	}
	n.AddVolumeCapabilityAccessModes(vcam)

	// SDFS plugin does not support ControllerServiceCapability now.
	// If support is added, it should set to appropriate
	// ControllerServiceCapability RPC types.
	n.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
	})

	n.AddNodeServiceCapabilities([]csi.NodeServiceCapability_RPC_Type{
		csi.NodeServiceCapability_RPC_GET_VOLUME_STATS,
		csi.NodeServiceCapability_RPC_UNKNOWN,
	})
	n.volumeLocks = NewVolumeLocks()
	return n
}

func NewNodeServer(n *Driver, mounter mount.Interface) *NodeServer {
	return &NodeServer{
		Driver:  n,
		mounter: mounter,
	}
}

func (n *Driver) Run(testMode bool) {
	versionMeta, err := GetVersionYAML()
	if err != nil {
		klog.Fatalf("%v", err)
	}
	klog.Infof("\nDRIVER INFORMATION:\n-------------------\n%s\n\nStreaming logs below:", versionMeta)

	n.ns = NewNodeServer(n, mount.New(""))

	s := NewNonBlockingGRPCServer()
	s.Start(n.endpoint,
		NewDefaultIdentityServer(n),
		// SDFS plugin has not implemented ControllerServer
		// using default controllerserver.
		NewControllerServer(n),
		n.ns,
		testMode)
	s.Wait()
}

func (n *Driver) AddVolumeCapabilityAccessModes(vc []csi.VolumeCapability_AccessMode_Mode) []*csi.VolumeCapability_AccessMode {
	var vca []*csi.VolumeCapability_AccessMode
	for _, c := range vc {
		klog.Infof("Enabling volume access mode: %v", c.String())
		vca = append(vca, &csi.VolumeCapability_AccessMode{Mode: c})
		n.cap[c] = true
	}
	return vca
}

func (n *Driver) AddControllerServiceCapabilities(cl []csi.ControllerServiceCapability_RPC_Type) {
	var csc []*csi.ControllerServiceCapability

	for _, c := range cl {
		klog.Infof("Enabling controller service capability: %v", c.String())
		csc = append(csc, NewControllerServiceCapability(c))
	}

	n.cscap = csc
}

func (n *Driver) AddNodeServiceCapabilities(nl []csi.NodeServiceCapability_RPC_Type) {
	var nsc []*csi.NodeServiceCapability
	for _, n := range nl {
		klog.Infof("Enabling node service capability: %v", n.String())
		nsc = append(nsc, NewNodeServiceCapability(n))
	}
	n.nscap = nsc
}

func IsCorruptedDir(dir string) bool {
	_, pathErr := mount.PathExists(dir)
	fmt.Printf("IsCorruptedDir(%s) returned with error: %v", dir, pathErr)
	return pathErr != nil && mount.IsCorruptedMnt(pathErr)
}
