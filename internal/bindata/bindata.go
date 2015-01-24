package bindata

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templs_index_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x91\x31\x53\xe3\x30\x10\x85\xfb\xfc\x8a\xb5\xeb\xb3\x35\xb9\xdc\xcd\x1d\x8c\xec\x06\x32\x81\x0a\x0a\x1a\x26\x93\x42\x91\xd7\x91\x40\xb6\x3c\xda\x4d\x86\xf0\xeb\x91\x62\x93\x14\x50\xd0\x48\xb3\x6f\xdf\xbe\x6f\xa4\x95\xd9\xed\xc3\xcd\xd3\xf3\xe3\x12\x0c\x77\xae\x9e\xc9\x74\x81\x53\xfd\xae\xca\xb1\xcf\x93\x80\xaa\xa9\x67\x00\xb2\x43\x56\xa0\x8d\x0a\x84\x5c\xe5\x7b\x6e\x8b\xff\xb1\x9f\x3a\x6c\xd9\x61\xbd\xf2\x69\x0c\xc8\xbe\xa3\x6f\x81\xed\x40\x52\x8c\x9d\x93\xc9\xd9\xfe\x15\x02\xba\x2a\x27\xe3\x03\xeb\x3d\x83\xd5\xbe\xcf\xc1\x04\x6c\xab\x5c\xd8\x4e\xed\x90\x44\xab\x0e\x49\x2e\xe3\x91\x7f\x19\xe4\xa3\x43\x32\x88\xfc\x39\x65\x98\x07\xba\x16\xa2\x53\x6f\xba\xe9\xcb\xad\xf7\x4c\x1c\xd4\x90\x0a\xed\x3b\x71\x16\xc4\xa2\x5c\x94\xbf\x85\x26\xba\x68\x65\x67\xa3\x8b\x68\xe2\x64\x45\xb1\xb6\x2d\x38\x86\xfb\x25\x5c\x6d\xd2\x9b\xa3\x4a\x3a\xd8\x81\x81\x82\xbe\xd0\x3c\x51\x39\x11\x13\x24\x7d\xd9\x5f\x32\xf6\x10\x21\xff\x22\xe4\x5c\x9f\x00\x2f\x31\x5f\x8a\x31\xe6\xe7\x99\x01\x69\xf0\x7d\x23\xe6\xe5\x9f\x98\x38\x55\xdf\xe6\xc9\x6c\x8d\x7d\x63\xdb\x4d\x51\xc4\x6d\x89\x71\x5d\x72\xeb\x9b\xe3\x44\x33\xf3\xfa\x0e\x9d\xf3\xbf\x60\xa5\x9c\xca\xa2\x65\x9e\x8c\xa3\x43\x8a\x71\xef\x1f\x01\x00\x00\xff\xff\xf6\x3e\x5c\xe3\x08\x02\x00\x00")

func templs_index_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templs_index_tmpl,
		"templs/index.tmpl",
	)
}

func templs_index_tmpl() (*asset, error) {
	bytes, err := templs_index_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templs/index.tmpl", size: 520, mode: os.FileMode(420), modTime: time.Unix(1422094945, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pub_images_favicon_ico = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x62\x60\x60\x04\x42\x01\x01\x10\xad\xc0\x90\xc1\xc2\xc0\x20\xc6\xc0\xc0\xa0\x01\xc4\x40\x21\xa0\x08\x44\x1c\x0c\x80\x72\x42\xdc\x10\x0c\x03\x07\x4b\xb3\x28\xc6\x5a\xca\xcc\xff\xb1\x61\x7c\x72\x30\x79\x98\x1a\x05\x69\x26\x0c\x0c\x53\x83\x4f\x0e\x59\x3f\xb2\x9b\xb0\xa9\x43\x96\x23\xa4\x1f\x9f\x19\x84\xf4\x23\xf3\xb1\x99\x81\x4f\x3f\x72\x58\xa0\xbb\x03\x97\x79\xc8\x61\x8d\xae\x97\xd8\x30\x21\x55\x2f\xae\xf8\x43\x77\x2f\x36\xbd\xc4\x86\x1f\xb2\x79\xc4\x86\x1f\x31\xfe\x21\x36\xfe\x71\x99\x81\x4b\x3f\x3a\x26\x26\x6d\x13\x93\x47\x08\xe5\x1f\x4a\xf0\xff\xff\x0c\x0c\x8d\x8c\x0c\x0c\x0d\x8d\x0c\x0c\x73\x2a\x19\x18\xe6\x59\x32\x30\xcc\x94\x64\x60\x98\x38\x91\x81\x61\xf1\x63\x06\x86\xe3\x47\x19\x18\x3a\x3b\x19\x18\x66\xcc\x44\xc8\xc3\xd4\x83\xf4\x02\x02\x00\x00\xff\xff\xc9\x1c\xf2\x68\x7e\x04\x00\x00")

func pub_images_favicon_ico_bytes() ([]byte, error) {
	return bindata_read(
		_pub_images_favicon_ico,
		"pub/images/favicon.ico",
	)
}

func pub_images_favicon_ico() (*asset, error) {
	bytes, err := pub_images_favicon_ico_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "pub/images/favicon.ico", size: 1150, mode: os.FileMode(420), modTime: time.Unix(1422094734, 0)}
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
	"templs/index.tmpl":      templs_index_tmpl,
	"pub/images/favicon.ico": pub_images_favicon_ico,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"pub": &_bintree_t{nil, map[string]*_bintree_t{
		"images": &_bintree_t{nil, map[string]*_bintree_t{
			"favicon.ico": &_bintree_t{pub_images_favicon_ico, map[string]*_bintree_t{}},
		}},
	}},
	"templs": &_bintree_t{nil, map[string]*_bintree_t{
		"index.tmpl": &_bintree_t{templs_index_tmpl, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
