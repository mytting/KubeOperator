// Code generated by go-bindata. (@generated) DO NOT EDIT.

//Package i18n generated by go-bindata.// sources:
// locales/en-US/home.yml
// locales/zh-CN/home.yml
package i18n

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// ModTime return file modify time
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

var _localesEnUsHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\xcd\x6e\xeb\x36\x13\xdd\xf3\x29\x26\xf2\xf6\xc3\xf7\x00\xde\x31\x12\x93\xa8\x57\x96\x04\xfd\xdc\x36\xdd\x08\xb4\x34\xb6\xd9\xc8\xa4\x40\x52\x4d\xd3\x5d\xdf\xab\xef\xd4\x57\x28\x28\xd3\xb2\x72\xe3\x20\x09\xba\x0a\x02\xe8\xcc\x9c\x33\x67\xe6\xd0\xab\x56\x1d\x8f\x4a\x92\x94\x6e\x58\xc3\x7e\x89\xcb\xaa\x5c\x43\x90\xf2\x23\x02\xef\x35\xf2\xee\x05\xf0\x0f\x61\xac\x09\x08\x59\xb5\xfd\x68\x2c\x6a\x12\x26\x75\x59\xb1\xa2\x89\x58\xc2\x2a\xd6\xdc\xd1\x38\x61\xd1\x1a\x82\x96\x4b\x90\xca\x42\x87\x3d\x5a\x04\xff\x39\x08\x09\xed\xa8\x35\x4a\x0b\xc6\x72\x8b\x01\x59\xb5\x1a\x3b\x94\x56\xf0\x9e\xbc\x2a\xd2\x14\xac\xcc\xea\x22\x64\x6b\x08\xee\xb8\xe8\xb1\x03\xab\x7c\xbd\x1b\xa8\x0e\xa8\x11\x84\x01\x2e\x81\x1b\xa3\x5a\xc1\x2d\x76\x70\x50\xc6\xc2\x28\x3b\xd4\x60\x0f\xc2\xc0\x13\xbe\x04\xef\x94\x6d\x7e\xcd\xd2\x2f\xd5\xfe\x53\x49\xbc\x52\xfb\x8e\xd6\x49\xd5\x84\x05\x8b\x58\x5a\xc5\x34\x69\x42\x9a\x36\x69\x56\xf9\x91\xac\x21\x88\x70\xc7\xc7\xde\xc2\x45\x29\xb4\x5c\xba\xe9\x6c\xd1\x37\xed\xdc\x4c\x1d\x79\xf2\x90\x95\x55\x43\x93\x82\xd1\xe8\xf1\xe2\xc2\x83\xd3\xf5\xa3\x0b\x5e\xd7\x84\x98\x07\x7f\x55\xce\x69\x2e\x4e\x91\x2f\xb1\x90\xf5\x2c\xec\x01\xec\x61\xf6\xe8\x5a\xdd\xe6\xf6\xb1\xc9\x8b\xec\x27\x16\x56\xff\xa9\xc5\xa0\xd5\x6f\xd8\xda\x80\x94\x8f\x65\xc5\x36\x4d\x9c\x4f\x93\xba\xcb\xea\xd4\x71\xcf\x7b\xe4\x06\x61\x27\xfa\xde\x6d\x8a\x43\xf4\xaa\xe5\x3d\xc4\x39\xec\x84\x36\xf6\x9f\xbf\xff\x0a\x48\x9a\x39\x1c\xfd\x4e\xe3\x84\xde\x26\x6e\xc0\xa9\x82\x78\x00\xfe\x3b\x17\x3d\xdf\xf6\xe8\x66\x39\x1a\xd4\x24\xa7\x65\xf9\x73\x56\x44\x53\x93\x0d\xad\xc2\x87\x35\x04\x8e\xeb\xc0\x8d\x79\x56\xba\x73\x7c\x85\x6c\x95\xd6\x13\xab\xac\x88\xef\xe3\x94\x26\x6f\xbe\x57\x5a\xec\x85\xe4\xfd\x7b\xc0\xba\x64\x45\x13\x97\x13\x8e\x86\x55\xfc\x9d\x79\xa0\xa3\xe1\xbe\x75\x66\xa3\x74\xe4\xba\x1b\xf0\x32\x5b\x25\x2d\x6f\xed\x24\x93\x77\x47\x21\x85\xb1\x9a\x5b\xa5\x6f\x7c\xc1\xe5\x68\x52\x05\x66\x6c\x0f\x53\xc1\x69\x0a\x34\xda\xc4\xe9\xdb\x5d\x73\x4d\x3b\xbf\x6f\x53\xd1\x13\x85\x37\xfb\x76\xf3\x9a\x74\xc1\x12\x5a\xb1\x68\x61\x72\xed\x60\x07\xee\xa8\x2f\xad\xf4\x0e\x5e\xa1\x50\xe7\x11\xfd\x1c\x05\xec\xc4\x89\x01\x21\x2b\x8d\x7b\xa1\xe4\x79\xe5\x0a\x76\x1f\x67\xe9\x67\x03\x00\x4e\xe0\x8f\x96\xce\xdd\xad\x5b\x09\xf7\xf7\xdc\xc8\xdd\xfe\xa7\xdb\x4c\x87\xff\xd1\x66\xf7\x5c\xba\x26\x6a\x40\x69\x2c\x6f\x9f\xc8\x3d\xab\xce\x7a\x58\x51\x64\xc5\xc9\x44\x4f\x79\xa7\x46\x39\xdd\xbc\x9f\xe7\x06\x8f\x5b\xd4\xb3\x25\x34\x8a\x96\x16\x6c\x11\x25\xf0\xae\xc3\x25\x64\x8e\x1e\xef\xd9\xfb\xb9\xe3\x01\xd7\x42\xe7\x8c\x7d\xa0\x65\xe3\xa3\xdc\xdd\xa1\x07\x2c\x84\xfa\x70\xb8\x81\xf0\xca\x26\x11\xb2\x92\xaa\x43\x92\x66\x11\x9b\xb3\xab\xa8\xd3\x34\x4e\xef\x9b\x8a\x96\xdf\xd6\x10\xd0\xae\x03\xf7\x11\x28\x7d\x7e\x14\xa6\x7f\xcf\x43\xd5\xa3\x94\x42\xee\xff\x37\x9c\x8e\xe3\x99\x0b\x0b\xc2\x42\xa7\x24\xfe\xdf\xa9\xde\xf2\xf6\x69\x1c\x68\xdb\xaa\x51\x5a\x92\xd3\x82\x6e\x1a\xb6\xc9\xab\xc7\x35\x04\xb1\x34\xe3\x6e\x27\x5a\xe1\xde\x95\x81\x6b\x7e\x44\x8b\xda\xb8\xa8\xa8\x9a\xb2\xce\xf3\xac\x98\x56\x5a\x9a\x71\x18\x94\x76\x7a\xec\xcb\x80\x01\x09\x1f\x58\xf8\xed\x92\x9d\xdf\x51\x8b\x9d\x68\xb9\x9d\x2c\x9a\x76\x61\x4e\xc3\x5b\x1a\x7e\xab\xf3\x86\x86\x61\x56\xa7\x5f\xc9\xc5\x57\xc4\x3f\x1d\x90\x64\xe5\x36\xea\x87\xa7\xeb\x13\xdd\x1c\xea\x0b\x4d\xbc\xad\xb7\x13\x47\xe2\x35\xde\xc5\x09\x3b\xc5\x42\x5a\x27\xc9\x25\x97\xbd\x6d\x76\x16\xe5\x92\x1a\x61\x8b\x3b\xa5\x11\xcc\xb3\xb0\xed\x41\xc8\xfd\xf2\x03\x7e\x92\x1d\xcc\xbf\x13\xe2\xd2\x4f\xd2\x07\xc5\xfc\xc3\xc0\x6d\xb9\x03\x3b\x20\x76\x30\x0e\x53\xc4\x2c\x60\x05\x2b\xab\xac\x60\x6f\x71\x1a\x8d\x55\x5a\xc8\xfd\x84\x98\xef\xa3\x40\xa3\x46\xdd\xe2\xdb\x19\x2e\x64\x7e\x28\xee\xf2\xec\x5c\x7f\x15\x2f\x47\x33\xbf\x81\xe7\xd1\x6f\xb1\x57\x72\x6f\x9c\x43\x8b\xe7\xd5\xd9\x24\x0c\xa8\x01\xf5\x69\xcf\x2e\x57\x39\xa0\xde\x29\x7d\xf4\x27\x45\xc8\xbf\x01\x00\x00\xff\xff\x87\x35\x9a\xce\x91\x09\x00\x00")

func localesEnUsHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesEnUsHomeYml,
		"locales/en-US/home.yml",
	)
}

func localesEnUsHomeYml() (*asset, error) {
	bytes, err := localesEnUsHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/en-US/home.yml", size: 2449, mode: os.FileMode(420), modTime: time.Unix(1597635710, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _localesZhCnHomeYml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x55\x4d\x53\xdb\x56\x14\xdd\xeb\x57\x78\xcc\xb6\xd3\x1f\xc0\xee\x21\x3d\x40\x45\x96\x34\x4f\x52\x5a\xba\xd1\x10\xa3\x69\x69\xc0\x62\x8c\xd9\x74\x55\x43\x00\x43\xac\x40\x13\x93\x42\xf1\x04\x4c\x4d\xf1\x24\xb1\xcd\x47\x8a\x8d\x85\xc9\x9f\xd1\x7b\x92\x57\xfc\x85\xce\x93\x6c\x21\xdb\xa5\xf1\x56\x73\xde\xb9\xe7\x9e\x7b\xee\xd5\x58\xd2\x5c\x5a\x32\x53\x8c\x08\x12\x50\x87\x3f\xf0\x8a\xaa\x8c\xc7\xe2\x78\xcf\x72\xcf\x2f\x70\xe3\x0a\x57\x0f\x70\xb1\x12\x67\x98\xb1\xe4\xe2\xea\x4a\xc6\x48\x33\xac\xa0\x29\x2a\x44\x3a\x07\x05\xa8\x42\x7d\x12\xf0\x02\xe4\xc6\x63\x71\xf2\xc7\x09\xb9\xde\xc7\xb9\x93\xce\x61\x19\xb7\xdf\xe2\x6d\xcb\xdd\xb9\x21\xbf\x65\xdd\x3f\x5f\x76\x8e\x36\xdd\xfb\xb2\x4f\x92\x36\xe6\x8d\x54\x66\x61\x6e\x91\xe9\x7b\xaf\x23\xa8\x48\x1a\x62\x21\xad\x1d\x50\x94\x2f\xbd\xcf\x67\x0f\x77\x59\xaf\x7e\xe6\x9e\x1f\x74\xde\x9c\x39\xcd\x57\xa4\xb8\x8d\x37\xae\xbd\x6c\xc1\x69\xda\xa4\xd8\x8a\x3f\x41\xa2\xff\x28\x89\xa3\x32\xe1\xdd\xba\x5b\xa8\xe0\xbc\x4f\x36\x09\x34\x41\xd5\x59\x04\x39\x28\xaa\x3c\x10\x74\x16\x88\xba\x28\xa9\xdd\x66\xc7\x63\xf1\x8e\x7d\xe0\xd5\xca\x78\xab\x4a\xac\x9a\xd3\xb4\xbc\xf5\x76\x50\xe4\xe1\x2e\x4b\xfb\xfb\xd9\x5c\xc9\x30\xd3\x92\xa2\xea\x40\x40\x10\x70\xb3\x8f\x96\x06\x92\x23\x96\x76\xb5\xfb\xe8\xd0\xc5\x61\xc9\xe1\x3b\xd7\xde\x0d\x24\xf7\xec\x1c\x26\xd0\x27\x66\x75\x19\x49\xdf\x41\x56\x1d\x95\xab\x74\xeb\x1e\xd5\x7c\xf5\xca\xac\xa2\xc2\x84\xce\xcb\x7e\xc7\x93\x92\x26\x52\x41\x5e\xbd\x81\x37\x72\xb8\xf4\x11\x6f\x1e\x92\xe2\x27\x52\x6c\xf1\xb2\x0f\x17\x25\x0a\x05\xcf\x00\x2f\x80\x09\x81\x7a\xc3\xcb\x31\xef\x9f\x97\xa4\xb5\x47\x8d\xb9\xb9\xa6\x7e\xac\xae\x18\x69\x46\x06\x8a\xf2\xbd\x84\x38\x9f\x37\x01\x54\x76\x9a\x8a\xab\x6f\xba\x27\xd9\x4e\xe1\xd0\xab\xd7\xe3\x8c\x84\xf8\x29\x5e\x04\x42\x3f\xe4\xf5\x71\x3f\x4a\x53\x20\xd2\x79\xc5\x07\x01\x56\xe5\x9f\xd1\xaa\x6e\xa1\x42\x72\x0d\x52\xfc\x80\xf7\xe8\x24\xfd\x46\x1b\x5e\xb6\xe0\x5e\xdb\x6e\xad\xe4\xee\x6d\xe2\xdf\x0f\x7c\xc1\xfe\xeb\x68\x6b\x34\xb1\xd5\x72\xf0\xde\x47\x00\x2e\xc1\x8b\x4f\xcd\x3c\x36\x37\xbf\xb4\x90\x8a\x05\xf0\x60\xf4\xde\xe9\xc7\xc8\xf4\xa3\xea\x10\x14\x80\x0a\xb9\xc8\x30\xba\x32\xaf\x4a\x61\xf2\x02\xeb\x07\xab\x6a\x32\x07\xfc\xaa\x60\xa8\x9c\xf3\xa5\x46\x0a\xb7\x41\xd2\x98\xb1\xb4\xf1\xd3\x82\x99\xea\x85\x00\xc1\x29\x5e\x12\x47\xda\x26\x9c\x6f\xe1\xe3\xe3\x68\x08\x22\x3b\xc0\x8c\xfd\x6a\xa6\x8c\x1e\x2b\xdd\xa3\xd1\x38\x7b\x0c\x7d\xd9\x5a\xaf\xb8\xed\x2b\xaf\x56\xc2\xb9\x37\x94\xd9\x5c\x36\x52\x2b\x99\xb9\xe4\x0b\x66\x0a\xaa\x3d\xc5\x10\x21\x09\xd1\x61\x6c\xdf\x3b\x4d\x0b\xe7\x2e\x02\x79\x14\xbf\x9c\x36\x7f\x31\x92\x99\x84\xb1\xf4\xdc\x48\x87\xf6\x02\x8e\x0b\xed\x0c\xaa\x91\x86\x8d\x77\x4e\x22\x2f\xc2\x4d\xee\xda\xff\xd4\x48\x83\x09\x0c\xad\x71\xef\xd5\x34\x50\xf4\xee\xb5\xa3\x4f\x7c\x70\x74\x03\x1f\xee\xb2\xff\x71\x02\x52\xe6\xbc\xc1\x88\x12\x07\xc3\x13\x80\x34\x51\xe4\xc5\x29\x5d\x05\xca\x0c\x75\x6f\xe3\xc6\xb1\xdf\x79\x3b\x6b\xee\xda\x2d\xd9\xbf\xe8\x6c\xed\x92\xb7\x96\xd3\x2e\x92\xea\x5f\xb8\x58\x21\xdb\xe7\x5e\x29\xff\x4d\xcc\xab\x37\xdc\xea\x36\xbe\xdf\xc0\xb5\x75\xc7\xfe\x14\x7c\xc6\xb5\x3c\xa9\xef\x7f\x4b\xcb\x3c\x9f\x4b\xbe\x58\x5d\x06\xc9\xa4\xb9\x9a\xca\x30\x32\x40\x20\xa1\xc3\x84\xac\xce\xd2\x0a\xbb\x6b\x64\xff\xa2\xb7\x85\xb4\x71\x45\x93\x65\x09\xa9\xfe\x19\xb2\x48\xa1\x4e\xf2\xf4\x2e\xbb\x97\x36\x7e\xff\x2a\xce\xb0\xd3\x90\x9d\x89\x5c\xf2\x93\x52\xe7\x43\x3e\x1c\x6f\x78\x68\x26\x00\x3b\xa3\xc9\x3a\x60\x59\x49\x13\x47\x3d\x39\xb8\xbc\xe5\xd8\x6d\xef\xf3\xdf\x78\xb7\xf1\xc4\xe1\x61\xc6\x96\x17\xe7\x52\x03\xb7\xfc\x2b\xb4\xd1\x6c\x0d\xd3\x46\xfe\x57\x13\xbe\x4f\x4c\x57\xfb\x24\x2f\xc0\x60\x43\x45\x4d\x10\x1e\xcf\x5b\x97\x9f\x4a\x25\xef\xb6\x1c\xfb\x06\x6f\x5a\x38\xb7\x45\xac\xd3\xa8\xfe\x78\xf8\xf3\xe3\x95\xae\x1b\x34\x18\x7e\x18\x82\xe9\x05\x68\xa7\x59\xf5\x1b\x8b\xa0\x11\x54\x54\x09\xc1\x01\x38\xc9\x9e\xe2\xb2\xd5\x83\x87\xf9\x45\xc6\x8a\xb9\x9a\x4e\x1a\xc3\x96\x44\xda\xf8\x1f\xf1\xd1\xa1\x0d\xfc\x1d\x1e\xf3\xdc\xf7\x2f\xb8\x7c\xef\xb4\x5e\x0f\xa4\xda\xfb\x72\xe4\x95\xf2\xa4\x5a\x0e\xe2\xd9\xbd\x3b\xcc\xbf\x01\x00\x00\xff\xff\xe6\xb3\x85\x3f\x31\x08\x00\x00")

func localesZhCnHomeYmlBytes() ([]byte, error) {
	return bindataRead(
		_localesZhCnHomeYml,
		"locales/zh-CN/home.yml",
	)
}

func localesZhCnHomeYml() (*asset, error) {
	bytes, err := localesZhCnHomeYmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "locales/zh-CN/home.yml", size: 2097, mode: os.FileMode(420), modTime: time.Unix(1597635710, 0)}
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
	"locales/en-US/home.yml": localesEnUsHomeYml,
	"locales/zh-CN/home.yml": localesZhCnHomeYml,
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
	"locales": &bintree{nil, map[string]*bintree{
		"en-US": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesEnUsHomeYml, map[string]*bintree{}},
		}},
		"zh-CN": &bintree{nil, map[string]*bintree{
			"home.yml": &bintree{localesZhCnHomeYml, map[string]*bintree{}},
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
