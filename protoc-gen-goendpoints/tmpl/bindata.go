// Code generated by go-bindata.
// sources:
// bindata.go
// endpoints.pb.go.tmpl
// DO NOT EDIT!

package tmpl

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

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

var _TmplBindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func TmplBindataGoBytes() ([]byte, error) {
	return bindataRead(
		_TmplBindataGo,
		"../tmpl/bindata.go",
	)
}

func TmplBindataGo() (*asset, error) {
	bytes, err := TmplBindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../tmpl/bindata.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1534425729, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _TmplEndpointsPbGoTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x93\x61\x4b\xfb\x30\x10\xc6\xdf\xf7\x53\x1c\x63\xfc\x69\xff\x8c\x7e\x80\x81\xaf\x44\x41\xd0\x39\x70\xe0\xeb\x92\xdd\x66\x91\x26\xf1\x7a\xb1\x93\x98\xef\x2e\x49\xd3\x8e\xc9\x6c\x2b\x0c\x5f\x25\x69\x2e\xcf\x3d\xcf\xef\xa8\x2e\xc4\x6b\xb1\x47\xb0\x16\xf2\x75\xdc\x3b\x97\x24\x65\xa5\x15\x31\xcc\x84\x92\x8c\x07\x9e\x25\xd6\xc2\xbc\x46\x7a\x2f\x05\xd6\xb0\xbc\x82\xfc\xa9\x3b\x38\x67\x2d\x15\x72\x8f\xd0\x57\xf8\x82\xbe\xda\xdf\x97\xbb\xfe\x9c\xaf\x8a\x0a\xc3\x23\x98\x57\xc8\x2f\x6a\xdb\xea\x3d\xc4\xbd\x73\x49\xa7\x17\xef\x83\x5a\xd5\x5f\xf3\x87\x0e\x7e\xe3\xa7\xcf\x3d\xf2\x9a\x14\xab\xa8\x7b\x23\xb7\x5a\x95\x92\x6f\x8d\x14\xb0\x33\x52\xa4\x82\x0f\x10\x73\xe4\xd7\xed\xba\x00\xc2\x37\xf8\x7f\x46\xe5\x4e\x6a\xc3\x1b\xdf\xc2\xb9\x0c\xd2\x73\x25\x8f\x86\x8f\x35\x0b\x40\x22\x45\x99\xb5\x28\xb7\x9e\x5c\x6f\xef\x5b\xde\xce\x57\x0d\x35\x93\x11\x0c\x36\x01\x00\x18\xce\x3a\x1e\x73\x22\x89\xce\xde\x14\x83\xcf\x54\x68\x8d\x74\x01\x9f\x1b\x75\xaf\x1a\x24\x70\xee\xf7\x36\xfd\xe8\x60\x85\xcd\x04\xa7\xe9\x70\x9e\xac\x1d\xf4\x58\xdc\x36\x27\x21\x1b\x92\xf0\x6f\xfa\x8b\x0b\xd0\x59\x02\xe6\xe3\x80\x16\x1d\x1d\xdf\x31\x30\x1a\x6e\x1b\x08\xa6\xcd\x94\x59\x67\x03\xf3\xf9\x9b\xff\xe7\x14\x7f\xf3\x13\x8e\x23\x33\x6f\x2b\xd8\xc8\x12\x17\xc1\x9c\x2e\x5f\x01\x00\x00\xff\xff\x60\x40\x12\x18\xdb\x04\x00\x00")

func TmplEndpointsPbGoTmplBytes() ([]byte, error) {
	return bindataRead(
		_TmplEndpointsPbGoTmpl,
		"../tmpl/endpoints.pb.go.tmpl",
	)
}

func TmplEndpointsPbGoTmpl() (*asset, error) {
	bytes, err := TmplEndpointsPbGoTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "../tmpl/endpoints.pb.go.tmpl", size: 1243, mode: os.FileMode(420), modTime: time.Unix(1534425726, 0)}
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
	"../tmpl/bindata.go": TmplBindataGo,
	"../tmpl/endpoints.pb.go.tmpl": TmplEndpointsPbGoTmpl,
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
	"..": &bintree{nil, map[string]*bintree{
		"tmpl": &bintree{nil, map[string]*bintree{
			"bindata.go": &bintree{TmplBindataGo, map[string]*bintree{}},
			"endpoints.pb.go.tmpl": &bintree{TmplEndpointsPbGoTmpl, map[string]*bintree{}},
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
