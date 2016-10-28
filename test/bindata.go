// Code generated by go-bindata.
// sources:
// test/default.yml
// test/dev.yml
// test/local.yml
// DO NOT EDIT!

package test

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

var _testDefaultYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x64\xcb\x31\x0e\x42\x21\x0c\x80\xe1\xbd\xa7\xe8\x05\x1c\x80\xad\x9b\x67\xf0\x04\x35\xaf\x46\x92\x67\x4b\xa0\xc8\xf5\x25\xe0\x62\xec\xd4\xbf\xcd\x77\xc8\x83\xfb\xe9\x04\x38\xa7\x58\xcb\x9e\x4d\x09\xbf\x67\x00\xe5\x97\x34\x82\x0b\x5e\xf5\xa8\x32\xe6\x72\x73\x79\x0b\x40\xe1\xea\x99\xcf\x0d\x4d\x85\x30\x84\xb0\xc2\x87\x11\xc6\x18\x77\x3c\xab\xcc\x5f\x4a\x69\x65\xeb\xf7\x2d\xfe\xd4\x8f\xfc\x04\x00\x00\xff\xff\x56\xfa\x3a\x43\x96\x00\x00\x00")

func testDefaultYmlBytes() ([]byte, error) {
	return bindataRead(
		_testDefaultYml,
		"test/default.yml",
	)
}

func testDefaultYml() (*asset, error) {
	bytes, err := testDefaultYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/default.yml", size: 150, mode: os.FileMode(420), modTime: time.Unix(1477645916, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testDevYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\xcd\x4f\x49\xb5\xe2\x52\x00\x82\x82\xfc\xe2\xcc\x92\xcc\xfc\x3c\x2b\x85\x94\xd4\x32\x2e\xae\xbc\xc4\xdc\xd4\x62\x2b\x2e\x5d\x05\xaf\xc4\xe2\xfc\x3c\x20\xed\x92\x58\x96\x99\xc2\xc5\x55\x90\x58\x54\x92\x99\x98\x03\xd1\x53\x5c\x9a\x04\x61\x80\x40\x7e\x5e\xaa\x15\x88\xe0\x02\x04\x00\x00\xff\xff\x58\xc0\xba\xfe\x54\x00\x00\x00")

func testDevYmlBytes() ([]byte, error) {
	return bindataRead(
		_testDevYml,
		"test/dev.yml",
	)
}

func testDevYml() (*asset, error) {
	bytes, err := testDevYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/dev.yml", size: 84, mode: os.FileMode(420), modTime: time.Unix(1477645898, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _testLocalYml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xca\xc9\x4f\x4e\xcc\xb1\xe2\x52\x00\x82\x82\xfc\xe2\xcc\x92\xcc\xfc\x3c\x2b\x05\xb0\x20\x17\x57\x5e\x62\x6e\x6a\xb1\x15\x97\xae\x82\x57\x7e\x46\x1e\x90\x0a\xc9\xcf\xe5\x02\x04\x00\x00\xff\xff\x1f\xba\xca\x9b\x30\x00\x00\x00")

func testLocalYmlBytes() ([]byte, error) {
	return bindataRead(
		_testLocalYml,
		"test/local.yml",
	)
}

func testLocalYml() (*asset, error) {
	bytes, err := testLocalYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "test/local.yml", size: 48, mode: os.FileMode(420), modTime: time.Unix(1477645673, 0)}
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
	"test/default.yml": testDefaultYml,
	"test/dev.yml": testDevYml,
	"test/local.yml": testLocalYml,
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
	"test": &bintree{nil, map[string]*bintree{
		"default.yml": &bintree{testDefaultYml, map[string]*bintree{}},
		"dev.yml": &bintree{testDevYml, map[string]*bintree{}},
		"local.yml": &bintree{testLocalYml, map[string]*bintree{}},
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

