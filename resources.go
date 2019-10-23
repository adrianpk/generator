// Code generated for package main by go-bindata DO NOT EDIT. (@generated)
// sources:
// assets/templates/migration.tmpl
// assets/templates/model.tmpl
// assets/templates/server.tmpl
package main

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _assetsTemplatesMigrationTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x91\xcf\x6b\xea\x40\x10\xc7\xcf\x99\xbf\x62\x5e\xf0\x90\x3c\x34\x81\x77\x14\x3c\x3c\x6c\xe9\x55\xd0\x63\xa1\xae\x71\xdc\x2e\xee\x0f\x3b\x3b\x42\x4a\xc8\xff\x5e\x36\x15\x6a\x5b\x5a\x2d\xf4\x98\xef\x64\x3e\xfb\x65\x3e\x07\xd5\xec\x95\x26\x74\x46\xb3\x12\x13\x3c\xd4\xf5\xea\xd1\x44\x6c\x82\xa3\x88\x3b\x0e\x0e\x75\x98\x6c\x8c\xdf\x2a\x51\x60\xdc\x21\xb0\x60\x6e\x83\xce\x01\xea\x1a\xe7\x4c\x4a\xa8\xeb\x26\x58\x2d\xec\x91\x95\x5d\xa8\xd8\x28\x3b\x57\x91\x70\xd2\xf7\x2b\xb5\xb1\xe7\xf0\xdd\xd1\x37\x58\x38\xfc\xeb\x8c\x2e\xaf\x5a\x2e\x4a\x24\xe6\xc0\xd8\x41\x26\x2d\x4e\x67\xe8\xaa\x3b\x92\x55\x5b\x94\x00\x59\x94\x94\xac\x07\xc4\x2b\x6d\x29\x4a\xc8\x91\x97\x44\x58\x03\x64\x69\x96\xdf\xfb\x3c\x7d\x03\x64\x0f\xe3\xc4\xc3\x19\x4a\x5b\xdd\xb6\xd4\x14\x51\x4a\xc8\xcc\x6e\x48\xff\xcc\xd0\x1b\x9b\x9e\xca\x98\xe4\xc8\x3e\xa5\x90\xf5\x9f\x31\x5d\x87\xac\xbc\x26\x1c\xed\xe9\xf9\xdf\x18\x47\xf1\xc9\x2e\x85\x8d\xd7\xa9\x50\xf5\xdf\x0a\xf1\x5b\x95\xbe\x07\xc4\x28\x78\xaa\x7a\xf6\xf3\x6f\x96\x4c\xa5\xc8\x6f\x71\x40\x9c\x46\xde\x58\xe8\x07\x53\x37\x1c\x0e\x17\x3c\x71\xb0\x76\xa3\x9a\xfd\x07\x4d\x97\x37\xaf\x97\x94\x58\xdf\x2b\xfa\xe2\x16\xd3\x1f\x1b\x7b\x77\x81\x97\x00\x00\x00\xff\xff\x81\xdf\x2a\x26\xe8\x02\x00\x00")

func assetsTemplatesMigrationTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesMigrationTmpl,
		"assets/templates/migration.tmpl",
	)
}

func assetsTemplatesMigrationTmpl() (*asset, error) {
	bytes, err := assetsTemplatesMigrationTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/migration.tmpl", size: 744, mode: os.FileMode(436), modTime: time.Unix(1571782333, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesModelTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x93\xcb\x6e\xdb\x3c\x10\x85\xd7\xe2\x53\x0c\x84\x00\xb1\x7f\xe4\x97\xf6\x01\xba\x08\xdc\x2e\x02\x34\x45\x5a\xb7\x5d\x87\x22\x47\xf2\x44\x24\x47\xe5\x05\x8e\x6b\xe8\xdd\x0b\x5d\x92\x3a\x89\xdd\x6c\xb2\xb3\x39\x33\xdf\x1c\x1d\x1e\x76\x52\xb5\xb2\x41\xb0\xac\xd1\x08\x41\xb6\x63\x1f\x61\x21\xb2\xbc\xa1\xb8\x49\x55\xa1\xd8\x96\x71\x4b\xee\xbe\x4c\x89\x74\x2e\x32\x0b\x43\xc9\xc8\xa9\x64\xa9\xf5\xbc\xc5\xdf\x68\xca\x4a\xaa\x16\x9d\x2e\x47\x54\x2e\x96\x42\xc4\x5d\x87\x03\xab\x2c\x61\xbf\x2f\xd6\xe4\x9a\x64\xa4\xbf\x95\x41\x49\xb3\x92\x01\xfb\x7e\xde\x9b\x9d\x2a\x87\xe8\x93\x8a\xb0\x17\x59\x66\x8b\x6b\x8d\x2e\x52\x4d\x4a\x46\x62\x27\xb2\x6c\xbf\xff\x1f\xbc\x74\x0d\xc2\x59\x8b\xbb\x0b\x38\xeb\x3c\x77\x70\xf9\x01\x8a\x5b\xcf\xdd\x47\xac\x43\xdf\xcf\x6d\x54\x83\xe3\x38\x75\x14\xd7\xe1\x93\xad\x50\x6b\xd4\x73\x7d\x3a\xfe\x22\xed\xb0\xf3\xf1\xef\x5a\xd6\xf8\x7d\xd7\x0d\x47\x77\xba\xba\xcc\x9f\xce\xbf\x7e\x5e\xb1\x49\xd6\xf5\x7d\x0e\xf7\x81\xdd\x41\x69\xfe\x86\x95\xb4\x38\x7f\xc2\x05\x5b\x8a\x68\xbb\xb8\xcb\xef\x66\x2d\xe8\xf4\x93\xae\xe9\x37\x00\x80\x2d\xae\x92\xa6\x28\xb2\x7e\xb0\xae\x2c\x61\x8d\x71\xe5\x51\x46\xfc\x29\x4d\xc2\x00\x35\xfb\xc9\xae\x42\xd4\xc9\x29\x58\x1c\x98\x76\xb0\x10\xfe\x3b\x61\xe6\xf2\x25\x71\xb1\x04\xf4\x9e\x3d\xec\x05\x40\x59\xc2\x37\xec\x8c\x54\x08\xe7\xc7\xc1\xa3\x3f\xc5\x3a\x7a\x72\xcd\xf9\xa8\x46\x3a\x8e\x1b\xf4\x90\x1c\xfd\x4a\x08\x34\xdf\x0f\xfa\x89\x77\x5d\xc3\x30\x02\x9a\x31\x8c\xee\xe3\x03\x85\x08\xec\x61\xc7\x09\x3a\x8f\x35\xfe\x65\xb0\xc3\x62\x1e\x8b\xb0\x25\x63\xa0\x42\x48\x01\xf5\xb8\x69\x6d\x52\x03\x0d\x3a\xf4\xe3\xdd\xcf\x9d\x57\xd0\x19\x49\x6e\x48\x09\xb9\x66\x9a\x92\x26\x30\x6c\xd9\xb7\x50\xa5\x08\x34\xb3\xac\x6c\x11\x0c\x73\x0b\x81\x2c\x19\xe9\x41\x1a\x03\xc1\xa4\x66\xf2\x15\xa5\xda\x00\xb9\x10\xa5\x53\x38\xc1\xb9\x86\xb8\xa1\x00\x1e\x03\x27\xaf\xb0\x18\x73\x4c\x05\x16\x97\x70\xdc\xa0\xdb\xfa\x61\x08\x5f\x3e\x60\xbb\xfa\x21\x7f\x16\xeb\xd7\x7d\x6f\xbb\x7c\x0a\xf0\xe2\x25\x14\x2f\xef\xf5\xe4\xda\xe5\x49\xe2\x98\xbd\x57\xa0\xa5\xc8\x3c\xc6\xe4\x1d\x38\x32\xa2\x7f\x8c\xe5\x8f\x4e\xbf\x73\x2c\x0f\x89\x07\xb1\x7c\x53\xed\xf3\xb9\x23\x6a\x6f\x64\x54\x1b\x50\xec\x34\x0d\x56\xbd\x83\xda\x91\xb8\x88\xea\x1f\x2d\x15\xb3\x19\xd4\xfb\xf9\x9a\x6f\x86\x85\xe3\xdc\x8a\xa7\xb7\x3f\xeb\xf4\xa2\x17\x7f\x02\x00\x00\xff\xff\x52\x01\xd0\xff\x82\x05\x00\x00")

func assetsTemplatesModelTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesModelTmpl,
		"assets/templates/model.tmpl",
	)
}

func assetsTemplatesModelTmpl() (*asset, error) {
	bytes, err := assetsTemplatesModelTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/model.tmpl", size: 1410, mode: os.FileMode(436), modTime: time.Unix(1571788688, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesServerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x41\x4f\xdc\x3c\x10\x3d\x6f\x7e\xc5\x28\x97\x2f\x41\xf9\x92\x7b\x25\x0e\x68\xab\x42\x0b\x6a\xa3\x85\x2d\x67\xd7\x99\x10\xab\x89\x1d\xc6\x63\xd8\x28\xca\x7f\xaf\x6c\x03\x0a\x55\x97\xd2\xc3\x4a\xeb\xd1\xbc\x37\xef\x3d\x8f\x53\x55\x70\x36\x8e\xa8\x1b\xe0\x4e\x59\x60\x03\x16\xe9\x01\x09\x5a\xd5\x63\x99\x54\x15\xec\x2d\x82\xb0\x20\x80\xb0\x45\x42\x2d\xb1\x00\xc5\x00\x83\x98\x80\xf0\xde\x29\x42\xb0\x66\x40\xc0\x46\x71\x99\x04\x88\x96\x66\x18\x50\x73\x24\x55\x2d\x48\xa3\x19\x0f\x7c\x89\x13\x28\x0b\xda\x30\x88\x9e\x50\x34\x13\x34\xd8\x2a\x8d\x4d\x18\xc5\xd3\x88\xeb\x56\xcb\xa4\xf4\x5d\x92\x48\xa3\x2d\x43\x96\x6c\xe6\x19\xca\x6b\xa5\xef\x5c\x2f\xe8\xca\x3c\x22\x49\x61\x11\x96\x65\xcb\x07\xdf\xbf\x82\x9e\x42\x7a\xac\x39\x4d\xf2\xa0\x72\x87\x63\x2f\x24\x42\x26\xe1\x64\x6b\x34\x0b\xa5\x91\x72\xf8\x31\xc1\x64\x1c\x81\x74\x96\xcd\xe0\x35\x38\xc9\xbe\xdf\x90\x1f\xf0\x80\xe4\x6d\x21\x0c\xc8\x9d\x69\x40\x69\x36\x20\xa0\x75\x92\x95\xd1\xde\xeb\x64\x1c\x34\x46\xff\xc7\xa0\x11\x1b\x1f\xa9\x60\x16\xb2\x03\x1f\x4f\xeb\xb4\xfc\x7d\xe2\x20\x7e\xe2\x99\x94\xc6\x69\x3e\xab\x3f\xef\x8c\x63\xa4\x6c\x14\xe4\x03\x94\x9d\x2a\x63\x25\x5f\xfd\x87\x39\xd9\x10\xb2\x23\x0d\xb1\x2f\xd6\xb3\xb4\xf2\xa6\xeb\xde\x91\xe8\x5f\x59\x2e\xc0\x0f\xce\x64\xf7\x8a\x70\x4e\x36\x1b\xd9\x95\xb5\xb1\x9c\xa5\x55\x5a\x80\x2c\xb7\x84\x82\x71\x1d\x5d\x2d\xac\x14\xfd\x36\x12\x7d\xb9\xfe\xf6\x35\x8f\xa8\x73\x7c\x01\x9d\x23\xbf\x07\x11\x35\xce\x73\x3a\xa7\xcb\x72\xec\x76\xe6\x39\x5d\xd2\x65\x79\x12\x6c\xff\xa0\x78\x63\x65\x57\xee\xad\x67\x3a\xba\x0d\xf9\x73\xdf\x3f\xaa\x0c\x98\xda\xbd\x60\xf6\x63\xf3\x9e\x38\x02\xec\x23\xf6\x18\xee\x20\x20\xe3\xe9\xaf\xc8\x25\x4f\xfc\x6f\x49\xe2\x66\xbc\x61\x29\xd3\x78\x60\xe8\x98\xc7\xf2\x42\xe8\xa6\xf7\x71\xac\x4f\xab\x9d\x58\x97\x3f\xf9\x18\x43\x96\x8f\xb1\xbe\x43\x3b\x1a\x6d\xf1\x96\x14\x23\x15\x40\x70\xf2\x54\xbf\x77\x68\x39\x66\x6c\x7b\x77\x07\x1f\x4e\x43\xf6\xfb\xdd\x55\x2d\x48\x0c\x19\x15\xc7\x1f\xd5\xff\x1e\x91\x86\x8b\xe6\x43\x40\xc6\xb7\x58\xde\x2a\xee\xbe\x8b\xde\x61\x46\xe5\x36\xd6\xb2\xbc\x78\xcb\xe8\x25\x4e\x05\x78\x3a\xcf\xe6\x4d\x97\xd7\xfe\x93\x74\x71\x73\x53\x67\x8f\x05\x50\xa0\x7c\xa6\x92\x7c\xc8\x9f\x23\xfc\x15\x00\x00\xff\xff\x20\x30\x9b\x85\xcc\x04\x00\x00")

func assetsTemplatesServerTmplBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesServerTmpl,
		"assets/templates/server.tmpl",
	)
}

func assetsTemplatesServerTmpl() (*asset, error) {
	bytes, err := assetsTemplatesServerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/server.tmpl", size: 1228, mode: os.FileMode(436), modTime: time.Unix(1571839273, 0)}
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
	"assets/templates/migration.tmpl": assetsTemplatesMigrationTmpl,
	"assets/templates/model.tmpl":     assetsTemplatesModelTmpl,
	"assets/templates/server.tmpl":    assetsTemplatesServerTmpl,
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
	"assets": &bintree{nil, map[string]*bintree{
		"templates": &bintree{nil, map[string]*bintree{
			"migration.tmpl": &bintree{assetsTemplatesMigrationTmpl, map[string]*bintree{}},
			"model.tmpl":     &bintree{assetsTemplatesModelTmpl, map[string]*bintree{}},
			"server.tmpl":    &bintree{assetsTemplatesServerTmpl, map[string]*bintree{}},
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