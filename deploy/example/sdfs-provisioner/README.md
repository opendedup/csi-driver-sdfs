# Set up a SDFS Server on a Kubernetes cluster

After the SDFS CSI Driver is deployed in your cluster, you can follow this documentation to quickly deploy some example applications. You can use SDFS CSI Driver to provision Persistent Volumes statically or dynamically. Please read Kubernetes Persistent Volumes for more information about Static and Dynamic provisioning.

There are multiple different SDFS server backends you can use for testing of 
the plugin
- To create a SDFS provisioner on your Kubernetes cluster, run the following command.

```bash
kubectl create -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-sdfs/master/deploy/example/sdfs-provisioner/sdfs-server.yaml
```

- During the deployment, a new service `sdfs-server` will be created which exposes the SDFS server endpoint `sdfs-server.default.svc.cluster.local` and the share path `/`. You can specify `PersistentVolume` or `StorageClass` using these information.

- Deploy the SDFS CSI driver, please refer to [install SDFS CSI driver](../../../docs/install-csi-driver.md).

- To check if the SDFS server is working, we can statically create a PersistentVolume and a PersistentVolumeClaim, and mount it onto a sample pod:

```bash
kubectl create -f https://raw.githubusercontent.com/kubernetes-csi/csi-driver-sdfs/master/deploy/example/sdfs-provisioner/nginx-pod.yaml
```

 - Verify if the SDFS server is functional, you can check the mount point from the example pod.

 ```bash
kubectl exec nginx-sdfs-example -- bash -c "findmnt /var/www -o TARGET,SOURCE,FSTYPE"
```

 - The output should look like the following:

 ```bash
TARGET   SOURCE                                 FSTYPE
/var/www sdfs-server.default.svc.cluster.local:/ sdfs4
```

```console
kubectl create secret generic sdfscreds --from-literal username=admin --from-literal password="PASSWORD"
```
