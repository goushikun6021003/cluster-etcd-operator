// Code generated by go-bindata.
// sources:
// bindata/v3.11.0/etcd/cm.yaml
// bindata/v3.11.0/etcd/defaultconfig.yaml
// bindata/v3.11.0/etcd/ns.yaml
// bindata/v3.11.0/etcd/operator-config.yaml
// bindata/v3.11.0/etcd/pod-cm.yaml
// bindata/v3.11.0/etcd/pod.yaml
// bindata/v3.11.0/etcd/sa.yaml
// bindata/v3.11.0/etcd/svc.yaml
// DO NOT EDIT!

package v311_00_assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _v3110EtcdCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-etcd
  name: config
data:
  config.yaml:
`)

func v3110EtcdCmYamlBytes() ([]byte, error) {
	return _v3110EtcdCmYaml, nil
}

func v3110EtcdCmYaml() (*asset, error) {
	bytes, err := v3110EtcdCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdDefaultconfigYaml = []byte(`apiVersion: kubecontrolplane.config.openshift.io/v1
kind: EtcdConfig
`)

func v3110EtcdDefaultconfigYamlBytes() ([]byte, error) {
	return _v3110EtcdDefaultconfigYaml, nil
}

func v3110EtcdDefaultconfigYaml() (*asset, error) {
	bytes, err := v3110EtcdDefaultconfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/defaultconfig.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdNsYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-etcd
  labels:
    openshift.io/run-level: "0"`)

func v3110EtcdNsYamlBytes() ([]byte, error) {
	return _v3110EtcdNsYaml, nil
}

func v3110EtcdNsYaml() (*asset, error) {
	bytes, err := v3110EtcdNsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/ns.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdOperatorConfigYaml = []byte(`apiVersion: operator.openshift.io/v1
kind: Etcd
metadata:
  name: cluster
spec:
  managementState: Managed
  # TODO this clearly needs to be fixed
  imagePullSpec: openshift/origin-hypershift:latest
`)

func v3110EtcdOperatorConfigYamlBytes() ([]byte, error) {
	return _v3110EtcdOperatorConfigYaml, nil
}

func v3110EtcdOperatorConfigYaml() (*asset, error) {
	bytes, err := v3110EtcdOperatorConfigYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/operator-config.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdPodCmYaml = []byte(`apiVersion: v1
kind: ConfigMap
metadata:
  namespace: openshift-etcd
  name: etcd-pod
data:
  pod.yaml:
  forceRedeploymentReason:
  version:
`)

func v3110EtcdPodCmYamlBytes() ([]byte, error) {
	return _v3110EtcdPodCmYaml, nil
}

func v3110EtcdPodCmYaml() (*asset, error) {
	bytes, err := v3110EtcdPodCmYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/pod-cm.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdPodYaml = []byte(`apiVersion: v1
kind: Pod
metadata:
  name: openshift-etcd
  namespace: openshift-etcd
  labels:
    app: etcd
    etcd: "true"
    revision: "REVISION"
spec:
  containers:
  - name: openshift-etcd
    image: ${IMAGE}
    imagePullPolicy: Always
    terminationMessagePolicy: FallbackToLogsOnError
    command: ["sleep"]
    args:
    - "7200"
    resources:
      requests:
        memory: 200Mi
        cpu: 100m
    ports:
      - containerPort: 10257
    volumeMounts:
    - mountPath: /etc/kubernetes/static-pod-resources
      name: resource-dir
#    livenessProbe:
#      httpGet:
#        scheme: HTTPS
#        port: 10257
#        path: healthz
#      initialDelaySeconds: 45
#      timeoutSeconds: 10
#    readinessProbe:
#      httpGet:
#        scheme: HTTPS
#        port: 10257
#        path: healthz
#      initialDelaySeconds: 10
#      timeoutSeconds: 10
  hostNetwork: true
  priorityClassName: system-node-critical
  tolerations:
  - operator: "Exists"
  volumes:
  - hostPath:
      path: /etc/kubernetes/static-pod-resources/etcd-pod-REVISION
    name: resource-dir

`)

func v3110EtcdPodYamlBytes() ([]byte, error) {
	return _v3110EtcdPodYaml, nil
}

func v3110EtcdPodYaml() (*asset, error) {
	bytes, err := v3110EtcdPodYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/pod.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdSaYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-etcd
  name: etcd-sa
`)

func v3110EtcdSaYamlBytes() ([]byte, error) {
	return _v3110EtcdSaYaml, nil
}

func v3110EtcdSaYaml() (*asset, error) {
	bytes, err := v3110EtcdSaYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/sa.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _v3110EtcdSvcYaml = []byte(`apiVersion: v1
kind: Service
metadata:
  namespace: openshift-etcd
  name: etcd
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: serving-cert
    prometheus.io/scrape: "true"
    prometheus.io/scheme: https
spec:
  selector:
    etcd: "true"
  ports:
  - name: https
    port: 443
    targetPort: 10257
`)

func v3110EtcdSvcYamlBytes() ([]byte, error) {
	return _v3110EtcdSvcYaml, nil
}

func v3110EtcdSvcYaml() (*asset, error) {
	bytes, err := v3110EtcdSvcYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "v3.11.0/etcd/svc.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"v3.11.0/etcd/cm.yaml":              v3110EtcdCmYaml,
	"v3.11.0/etcd/defaultconfig.yaml":   v3110EtcdDefaultconfigYaml,
	"v3.11.0/etcd/ns.yaml":              v3110EtcdNsYaml,
	"v3.11.0/etcd/operator-config.yaml": v3110EtcdOperatorConfigYaml,
	"v3.11.0/etcd/pod-cm.yaml":          v3110EtcdPodCmYaml,
	"v3.11.0/etcd/pod.yaml":             v3110EtcdPodYaml,
	"v3.11.0/etcd/sa.yaml":              v3110EtcdSaYaml,
	"v3.11.0/etcd/svc.yaml":             v3110EtcdSvcYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"v3.11.0": {nil, map[string]*bintree{
		"etcd": {nil, map[string]*bintree{
			"cm.yaml":              {v3110EtcdCmYaml, map[string]*bintree{}},
			"defaultconfig.yaml":   {v3110EtcdDefaultconfigYaml, map[string]*bintree{}},
			"ns.yaml":              {v3110EtcdNsYaml, map[string]*bintree{}},
			"operator-config.yaml": {v3110EtcdOperatorConfigYaml, map[string]*bintree{}},
			"pod-cm.yaml":          {v3110EtcdPodCmYaml, map[string]*bintree{}},
			"pod.yaml":             {v3110EtcdPodYaml, map[string]*bintree{}},
			"sa.yaml":              {v3110EtcdSaYaml, map[string]*bintree{}},
			"svc.yaml":             {v3110EtcdSvcYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
