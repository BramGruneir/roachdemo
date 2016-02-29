// Code generated by go-bindata.
// sources:
// assets/css/default.css
// assets/templates/cluster.html
// assets/templates/error.html
// assets/templates/layout.html
// assets/templates/log.html
// assets/templates/node.html
// assets/templates/notfound.html
// assets/templates/run.html
// DO NOT EDIT!

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

var _assetsCssDefaultCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4a\xca\x4f\xa9\x54\xa8\xe6\xe2\x2c\x48\x4c\x49\xc9\xcc\x4b\xd7\x2d\xc9\x2f\xb0\x52\x30\x37\x28\xa8\xb0\xe6\xaa\xe5\xca\x30\x02\x49\xe5\x26\x16\xa5\x67\xe6\x41\x64\x0c\xac\xe1\xfc\xa4\xfc\x92\x92\xfc\x5c\x2b\x05\x23\x88\x62\x40\x00\x00\x00\xff\xff\x28\xaf\x7a\xd0\x49\x00\x00\x00")

func assetsCssDefaultCssBytes() ([]byte, error) {
	return bindataRead(
		_assetsCssDefaultCss,
		"assets/css/default.css",
	)
}

func assetsCssDefaultCss() (*asset, error) {
	bytes, err := assetsCssDefaultCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/css/default.css", size: 73, mode: os.FileMode(420), modTime: time.Unix(1456760347, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesClusterHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x51\x6f\xdb\x36\x10\x7e\x4e\x7f\x05\xc7\x06\xb0\xfc\x20\xa9\x1d\xda\x61\xb0\x25\x03\x05\xf6\x32\xa0\x28\x86\x06\x7b\x1a\x86\x81\x16\x69\x89\x08\x4d\x72\x24\x15\x3b\x08\xfc\xdf\x77\x47\xca\xb2\x1c\x27\x4b\x5a\xa0\x08\x62\xfb\x74\xdf\x7d\x77\xfc\x78\x47\xaa\xf2\xe1\x5e\x89\xd5\x1b\x42\x02\x27\xdd\x07\xf2\x00\xbf\xb6\xcc\xb5\x52\x2f\xc8\xbb\x25\x18\x87\xe4\xb2\x4e\x44\xdf\x9a\x35\xb7\xad\x33\xbd\xe6\x0b\xa2\x8d\x16\x08\x59\x1b\xc7\x85\x3b\xd9\x31\xa4\x13\x8c\xc3\xe7\x45\xd0\xdb\xcd\x47\xfc\x3b\xe2\x8a\x2d\xdb\x77\x42\xb6\x5d\x18\x53\x98\x3b\xe1\x36\xca\xec\xf2\xfb\x05\xf1\x8d\x33\x4a\x2d\x63\x51\xfb\x3c\x01\x17\xe4\xd7\x77\x76\x3f\x12\x68\xc3\x45\x6e\xfa\x60\xfb\x30\xa9\x3e\x0f\xc6\x2e\xc8\xc7\x09\x2e\xb0\xb5\x12\x24\xb8\x45\x87\x09\x22\xb4\xe9\x9d\x37\x50\xb8\x35\x52\x07\xe1\x46\xa8\x65\x5a\x28\xf8\x72\xa6\x75\xc2\xfb\x0b\xda\x5f\x12\xed\x71\xdd\xef\xed\x9e\x78\xa3\x24\x27\x6f\x19\x63\x23\x8b\x32\xcd\xad\xe0\x31\xd8\x32\xce\xa5\x6e\x73\x25\x36\x58\x7e\x0a\x87\x2a\x82\x6c\x98\xca\x99\x92\x2d\xc8\x0d\xd4\xcb\x09\x36\x66\x1a\xa0\x8d\x51\x58\xe7\x19\x7d\x63\x74\x60\x52\x0f\x4b\x41\x79\x76\x92\x87\x0e\xd5\x99\xca\x73\x82\x15\xb8\x25\xc0\x4c\xba\x9f\x63\x08\x97\xde\x2a\x06\x22\x4b\xad\x00\x90\xaf\xb1\xe0\x14\x57\x95\x43\x5b\x54\xb0\x01\xd2\x86\xd5\x9b\xeb\x6c\xd3\xeb\x26\x48\xa3\xb3\x79\x8c\xbe\xce\xe8\x5f\x9c\x05\x06\x75\xb6\xad\x12\xf5\x2c\x18\xa3\x82\xb4\xb3\xbf\xe9\xbc\x18\x7e\x67\xf3\x65\x44\xce\x46\xf1\x67\xf3\xa2\x51\xb2\xb9\x3d\xb1\x89\x44\x47\x88\xdc\x90\xec\x3a\x13\x00\x75\xad\x08\xf3\x42\xfa\x8c\x32\x3a\x3f\xba\x09\x71\x22\xf4\x4e\x2f\xa3\x75\x88\x9f\x9d\x13\x1b\x52\x93\x69\x94\x65\x4e\xe8\xe0\xb3\x59\xcc\xb5\x91\x9a\x67\x14\xda\x17\x88\x0a\x16\x82\xcb\x66\x18\x33\x9b\x2f\xc7\x94\x91\xe3\xa7\x9a\x40\x73\x0a\x80\x0b\x7e\x4a\xb8\x83\x68\xb3\xc3\x7d\x64\x58\x6a\x31\xa4\xc3\xaf\x53\x15\x07\xe0\xc2\x7f\x90\x6c\x90\xaa\xe2\xf2\x8e\x34\x8a\x79\x5f\xd3\x51\x7d\x8a\x23\x56\x6d\x8c\xdb\x92\xad\x08\x9d\xe1\x35\xb5\xc6\x87\xf8\x18\x1c\x49\x9f\x21\x68\x10\x0b\x3f\xf3\xd4\x63\xd0\x46\xc9\x8c\xbd\x3b\x04\x61\x18\x6e\xe9\xd1\x42\xdb\x9d\x8c\xe8\x26\xb1\x27\x6a\xfa\x11\x5a\x82\xae\xbe\xc0\xa8\x54\x65\xe8\x9e\x01\xb1\x3e\x18\xba\xfa\xf3\xeb\xe7\xff\xc1\xbc\x4f\x4c\x9f\x4d\xeb\x5f\x46\x7d\x8a\x5b\xfc\x08\x08\xd6\x58\x25\x7a\x26\x2b\xa8\xc2\xda\xf0\xfb\x13\xf4\xe1\x81\x38\xa6\x5b\x41\x0a\xac\xdc\x93\xc3\x61\xea\xba\xc6\xc9\x27\x8b\x9a\x14\x53\x07\x68\x70\xd4\x11\x30\xb0\xc3\x05\x56\x71\x27\x00\x73\x66\x17\x7f\xb0\xde\x83\xae\x87\xc3\x8e\x39\x0d\x63\x01\x5e\xa1\x3c\xe2\x7c\xdf\x34\x30\xf6\xf8\x40\xf3\x14\x37\x78\x38\x56\xe3\x46\x07\x3d\x5f\x3e\x9f\x9a\xf0\x80\xc5\x56\xa9\x69\x89\x85\x96\x10\x55\x7c\x61\x5b\x11\xe3\x26\x46\x55\xb2\x33\x9a\xf2\x9c\xe7\x79\x5a\xe4\x80\xcd\x42\x3e\x92\xda\xbf\xa6\xff\xac\x15\xd3\xb7\x89\x3f\xf9\xbe\x95\xfe\xb1\x68\xe7\xb9\x27\xad\x1d\xcf\x5d\xd7\x6b\x7a\x1e\x1f\x0b\x1c\x20\xeb\xa0\x09\xfc\xe7\x7b\x1f\xbf\x60\xc0\x58\xaf\x02\x7d\x4e\x96\x12\xd8\xa2\x3d\xec\xd0\xef\xbf\xe1\x43\x1f\x38\x9c\xee\x74\x55\x79\x38\x96\x8f\xcc\xad\xba\xb7\x9d\x84\xe9\x22\xe3\xaf\x7c\x23\x95\x00\x58\x89\xb8\x15\x49\x61\x8f\x56\xff\x63\xca\x13\xce\x7d\x4f\x79\x10\x76\x51\x5e\x55\x82\xc2\x17\x1b\x32\x74\xdf\x39\x52\xc2\x38\x6b\x18\x67\x79\x09\x8f\xcd\xf9\xed\x9b\x2e\xfe\x25\xc5\x4d\x60\xa1\xf7\x84\xde\xc0\xf5\x63\x05\xa7\x17\x69\xd7\x7d\x08\xb0\x2e\x3c\xc9\x58\x1c\xef\xa7\x94\xf2\xd0\x8f\x20\xe4\xd3\x3a\x0f\xe3\x45\x57\x37\x88\xaa\xca\xc4\xf8\xba\x45\xbf\x32\xbb\xb1\xcf\x25\x4f\x13\x8c\xb9\x8d\x3d\xa5\xbe\x7a\x5a\x84\x74\x44\x7c\xa7\x06\xf0\xe6\xd0\x6f\xc5\x8b\x22\x7c\x8d\xb0\x1f\xa0\x82\xc5\xe2\x5f\x92\x21\xae\xf0\x52\x87\x8b\x0e\x7a\x4d\x5f\x4d\x0f\xf6\x29\xfe\x2a\x5e\x4b\x57\xb1\xed\xf0\x4d\x06\x07\xa0\xfe\xb0\x7a\x72\x15\xf0\xee\xf3\xa2\x62\x9f\x38\x27\xe9\x26\x1b\xca\x4e\x65\x5c\x3d\xba\x58\x4e\x57\x09\x18\x78\x79\xc6\x1b\xb8\xc4\x74\x70\x47\xa7\x39\xfb\x2f\x00\x00\xff\xff\xc4\xc7\x72\xa2\xfe\x0a\x00\x00")

func assetsTemplatesClusterHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesClusterHtml,
		"assets/templates/cluster.html",
	)
}

func assetsTemplatesClusterHtml() (*asset, error) {
	bytes, err := assetsTemplatesClusterHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/cluster.html", size: 2814, mode: os.FileMode(420), modTime: time.Unix(1456761227, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesErrorHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x4f\x6f\xdb\x3e\x0c\x3d\xff\xfa\x29\x54\xfd\xae\x93\x85\x74\x97\x1d\xec\x00\xdb\xd0\xc3\xee\x1b\xb0\x2b\x2d\xd1\xb6\x32\x59\x52\x25\x3a\x6d\x50\xf4\xbb\x4f\x8a\xff\x34\x2d\x3a\xa0\x18\x76\x08\x2c\x4a\xe4\x23\xf9\xf8\x98\xfa\x5a\x7b\x45\xa7\x80\x6c\xa0\xd1\xee\xaf\xea\xf2\x61\x16\x5c\xdf\x70\x74\x7c\x7f\xc5\x58\x3d\x20\xe8\x72\xc8\xc7\x11\x09\x98\x1a\x20\x26\xa4\x86\x4f\xd4\x89\x4f\xfc\xf2\x69\x20\x0a\x02\xef\x26\x73\x6c\xf8\x4f\xf1\xe3\xb3\xf8\xea\xc7\x00\x64\x5a\x8b\x9c\x29\xef\x08\x5d\x8e\xfb\x76\xdb\xa0\xee\xf1\x45\xa4\x83\x11\x1b\x7e\x34\x78\x1f\x7c\xa4\x0b\xe7\x7b\xa3\x69\x68\x34\x1e\x8d\x42\x71\x36\x3e\x30\xe3\x0c\x19\xb0\x22\x29\xb0\xd8\xec\x32\xd0\x8c\x44\x86\x2c\xee\x35\x8e\x9e\x61\x8c\x3e\xd6\x72\xbe\x59\x9e\xad\x71\xbf\x58\x44\xdb\xf0\x44\x27\x8b\x69\x40\xcc\x99\x86\x88\x5d\xc3\xa5\x54\xda\x1d\x52\xa5\xac\x9f\x74\x67\x21\x62\xa5\xfc\x28\xe1\x00\x0f\xd2\x9a\x36\x49\xba\x37\x44\x18\x45\xeb\x3d\x25\x8a\x10\xe4\xc7\x6a\x57\xed\xa4\x4a\x49\x6e\x77\x55\xb6\xb6\x6a\x92\x8a\x26\x10\x4b\x51\xbd\x03\xfe\x70\x37\x61\x3c\xc9\x9b\x33\xe6\x6c\x54\xa3\x71\xd5\x21\xe3\xd5\x72\x86\xda\xff\x05\xee\x9f\xca\x3e\x5c\x56\xfd\x32\xc9\x3b\xc8\x2a\x4d\x6b\xec\x60\xb2\xb4\xb4\x9c\x23\xe4\x2a\x94\xba\xf5\xfa\xb4\x14\xeb\xe0\xc8\x94\x85\x94\x1a\x9e\x8f\x2d\x44\x36\x7f\xc4\x12\xbe\x9a\x9d\x79\x40\x2d\xc8\x07\xce\xa2\xcf\x43\x2d\xde\xa6\xcf\xd2\xf1\x6e\xd1\x49\x06\xd3\x66\x03\x2b\xfa\x00\xe3\x72\x67\x9d\x9d\x8c\xce\x3e\xff\xd5\xd7\x42\xb0\x2f\x11\x9c\x66\xe5\x47\xbe\xef\x2d\xb2\x1e\x89\xf5\xd1\x4f\x01\x35\xeb\x7c\x64\x2d\x16\x3e\xd8\xe8\x5b\x93\x5f\xb5\x49\xc1\xc2\x89\x09\x51\x00\x2e\xf0\x97\xb2\x4a\x4b\x18\x0b\x7a\x69\x6b\x22\xf2\x8e\x95\x75\x69\xf8\x6c\xf0\x57\xfe\x73\x52\xce\x34\x10\x2c\x46\xa9\xd5\x5a\x08\x69\xbb\x86\xd8\x97\xf5\xf9\xbf\x4d\x02\x1f\x60\x0c\x16\xc5\x12\xbe\x7a\x8a\xdd\x9c\xb2\x4c\x3b\x80\x5b\x93\xa4\x28\xbc\xb3\x27\xbe\xff\x3e\xf7\xf6\xcc\x51\x9e\x5e\xf6\x7b\x2b\xc6\x64\xa6\x44\xc6\x3e\x4f\xf8\x5f\xfb\xd4\x72\xa6\x61\x36\xe0\x15\x19\x6d\x99\xc5\xa6\x19\x7e\x5e\xcc\x5a\x42\x61\x5a\x66\xaa\xb7\xb1\x3e\x1b\xb5\xcc\xa1\xab\x02\xdf\x9a\xf6\xb3\x16\x86\x9b\xfd\xe3\x63\x75\x5b\xd6\xfc\xe9\x29\x6b\xef\x66\x45\x58\xc0\x72\x69\x67\x15\xe6\xa7\xf3\x3f\xdb\xef\x00\x00\x00\xff\xff\xe4\xaf\x29\x5d\xea\x04\x00\x00")

func assetsTemplatesErrorHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesErrorHtml,
		"assets/templates/error.html",
	)
}

func assetsTemplatesErrorHtml() (*asset, error) {
	bytes, err := assetsTemplatesErrorHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/error.html", size: 1258, mode: os.FileMode(420), modTime: time.Unix(1456714931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesLayoutHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x56\x4d\x8f\xe4\x34\x10\x3d\xb3\xbf\xc2\xeb\xbd\xae\x63\x0d\x5c\x38\x24\x2d\xc1\x80\xc4\x5e\x16\x84\x06\x89\xab\x13\x57\x12\x37\x8e\x9d\xb1\x2b\x3d\xd3\x6a\xf5\x7f\xa7\x9c\xaf\x4e\xf7\xc2\xd0\x02\x0e\x2d\xbb\xec\xf2\xab\x7a\xaf\xca\x4e\xe7\xef\xb5\xaf\xf0\xd8\x03\x6b\xb1\xb3\xbb\x77\x79\x1a\x98\x55\xae\x29\x38\x38\xbe\x7b\xc7\x58\xde\x82\xd2\x69\x42\xd3\x0e\x50\xb1\xaa\x55\x21\x02\x16\x7c\xc0\x5a\x7c\xcb\xb7\x5b\x2d\x62\x2f\xe0\x79\x30\x87\x82\xff\x2e\x7e\xfb\x4e\x3c\xfa\xae\x57\x68\x4a\x0b\x9c\x55\xde\x21\x38\x3a\xf7\xe9\xc7\x02\x74\x03\x57\x27\x9d\xea\xa0\xe0\x07\x03\x2f\xbd\x0f\xb8\x71\x7e\x31\x1a\xdb\x42\xc3\xc1\x54\x20\x46\xe3\x23\x33\xce\xa0\x51\x56\xc4\x4a\x59\x28\x1e\x08\x68\x42\x42\x83\x16\x76\xa7\x53\xf6\x94\x26\xe7\x73\x2e\xa7\x95\x79\xdb\x1a\xf7\x07\x0b\x60\x0b\x1e\xf1\x68\x21\xb6\x00\x14\xa9\x0d\x50\x17\x5c\xca\x4a\xbb\x7d\xcc\x2a\xeb\x07\x5d\x5b\x15\x20\xab\x7c\x27\xd5\x5e\xbd\x4a\x6b\xca\x28\xf1\xc5\x20\x42\x10\xa5\xf7\x18\x31\xa8\x5e\x7e\x93\x3d\x64\x0f\xb2\x8a\x51\xae\x6b\x19\x59\x6b\x36\xb1\x0a\xa6\x47\x16\x43\x75\x07\xfc\xfe\x79\x80\x70\x94\x5f\x8f\x98\x93\x91\x75\xc6\x65\x7b\xc2\xcb\xe5\x04\xb5\xfb\x17\xb8\x7f\x97\xf6\x7e\x9b\xf5\x75\x90\x3b\xc4\x4a\xa4\x35\xd4\x6a\xb0\x38\x53\xa6\x13\x72\x69\x94\xbc\xf4\xfa\x38\x27\xeb\xd4\x81\x55\x56\xc5\x58\x70\x9a\x96\x2a\xb0\x69\x10\xf3\xf1\xc5\xac\xcd\x2b\x68\x81\xbe\xe7\x2c\x78\x2a\x6a\xf2\x36\x0d\xb5\x8e\x77\x73\x9f\x10\x98\x36\x2b\x58\xea\x0f\x65\x1c\x31\xab\xed\x60\x34\xf9\x7c\x95\xbf\x17\x82\x7d\x1f\x94\xd3\x2c\xfd\xd0\x37\x8d\x05\xd6\x00\xb2\x26\xf8\xa1\x07\xcd\x6a\x1f\x58\x09\x49\x0f\xd6\xf9\xd2\xd0\xae\x36\xb1\xb7\xea\xc8\x84\x48\x00\x1b\xfc\x39\xad\x44\x09\x42\x42\x4f\xb4\x06\x44\xef\x58\xba\x2e\x05\x9f\x0c\x7e\xe3\x3f\x05\xe5\x4c\x2b\x54\xb3\x91\x72\xb5\x56\xf5\x71\x5d\x56\xa1\x49\xd7\xe7\x43\x19\x05\xbc\xaa\xae\xb7\x20\xe6\xe3\x8b\xa7\x78\x98\x42\xa6\x6a\xf7\xca\x2d\x41\x62\x10\xde\xd9\x23\xdf\x3d\x4d\xdc\x2e\x1a\x51\xf5\xc8\xef\xaf\xce\x18\x52\x4a\x10\xf6\x58\xe1\xff\xdb\x27\x97\x93\x0c\x93\xa1\x6e\xc4\x28\x53\x2d\xd6\x9e\xe1\x3b\x0d\x9d\xcf\xa5\x4a\x4a\x4b\x92\x9a\x3a\x6d\xaa\xd9\x23\xb1\x86\x0a\x19\xb6\x23\x25\x96\x5a\x2f\x7e\x4c\xd5\xea\x68\x48\xb5\xf4\xb4\x15\x96\x37\x61\x2c\xe3\xa8\xad\x71\xcd\x97\x95\x5b\x34\x64\x37\x9a\x72\x66\x34\x95\xed\x1f\x35\xcf\x07\xbb\xe1\xb1\xa0\xd0\xb0\x94\xe4\x74\x62\xa6\x66\xd9\xa3\x1d\x62\xea\xa4\xf3\x79\x56\xcb\x9a\x69\x07\x9e\x59\xf6\x8b\x6a\x80\xf1\xcf\x5e\x43\xe4\xe4\xb1\x00\xaa\x0a\xcd\x01\xf8\xe9\x04\x4e\x9f\xcf\x3b\x52\x6c\x15\xa7\x9a\xe0\x92\x3e\x39\x5d\xdd\x4b\x2c\x72\x5d\x63\xcc\xa1\x13\xee\x36\x2e\xbb\x0d\xfc\x93\x89\xe8\xc3\x31\x85\xbe\x8d\x3c\xe3\x6d\x62\x3b\x42\x93\xb4\x31\xc2\x66\x9f\xe9\x21\xa6\x7d\x2a\xf3\xb6\xfa\x8d\x3d\xf6\x6d\x6a\x01\xb6\xce\x84\x56\xb1\x2d\xbd\x0a\x7a\x6d\x09\x76\x8b\x72\x37\x9b\x5f\x07\xf7\x26\xa1\xd9\xe7\x3f\x10\x92\x61\x70\xeb\x22\x41\x65\x9f\x7e\xb8\x8f\x66\x7a\x1f\x2e\x0c\x53\xa2\x1f\xbe\x80\xb9\x87\xe7\x35\x99\x9f\x07\xec\x07\xe4\x57\xa4\xaf\x99\x5d\x08\xdd\x91\x64\x4d\x8f\xd9\x75\x19\x9e\xd2\x47\xfd\xed\xcc\x72\x39\xd8\xcb\x65\x9c\xdf\xd8\x8b\x91\x4b\x6a\xfa\x69\x4a\x9f\xd4\xc7\xe9\xf2\xd1\xc9\xf1\xda\x8f\x2f\x3c\x3d\xf9\xe3\xbf\x86\x3f\x03\x00\x00\xff\xff\xce\x2e\x8e\x1e\x46\x08\x00\x00")

func assetsTemplatesLayoutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesLayoutHtml,
		"assets/templates/layout.html",
	)
}

func assetsTemplatesLayoutHtml() (*asset, error) {
	bytes, err := assetsTemplatesLayoutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/layout.html", size: 2118, mode: os.FileMode(420), modTime: time.Unix(1456761931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesLogHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x44\x8e\xcd\xaa\x83\x30\x10\x85\xf7\x3e\xc5\xe0\x5d\xdf\x08\x77\xe9\x4d\x5d\x75\x53\x28\x16\x4a\x5f\x20\x9a\x41\xa5\x36\x23\x31\x69\x11\xc9\xbb\x77\x46\xfa\xb3\xc8\xe2\x7c\x27\x27\x5f\xf4\x1c\x96\x11\xab\x0c\x40\xb5\xe4\x82\x19\x1c\x7a\x58\x39\x3e\x06\x1b\xfa\x12\x4c\x0c\xf4\xcf\x31\xf1\x99\x3c\x6e\x55\x63\xda\x6b\xe7\x29\x3a\x5b\x82\x23\x87\xd2\x37\xe4\x2d\xfa\x6f\x4e\x99\x2e\x5e\x4f\x6b\x3b\xdc\xa1\x1d\xcd\x3c\xef\xf2\x8f\x23\x17\xa5\xee\xff\xaa\x75\x05\x55\x93\x45\x55\x9b\x1b\x42\x4a\xf0\xf3\x26\xe7\xe8\xd4\x61\x2f\xe8\x17\x84\x5d\x96\x49\x2e\xe8\x82\x57\x32\xe6\xef\x6c\xeb\x23\x75\xa7\x18\xa6\x18\xb6\x52\x28\xbb\xd9\x59\x65\xcf\x00\x00\x00\xff\xff\x0c\x02\x65\x26\xdd\x00\x00\x00")

func assetsTemplatesLogHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesLogHtml,
		"assets/templates/log.html",
	)
}

func assetsTemplatesLogHtml() (*asset, error) {
	bytes, err := assetsTemplatesLogHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/log.html", size: 221, mode: os.FileMode(420), modTime: time.Unix(1456760158, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesNodeHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x56\x4b\x6f\xe3\x36\x10\x3e\xc7\xbf\x82\x50\x82\xda\x3e\x58\xca\x65\x2f\x8e\x2c\xa0\xd8\xf6\xb0\x40\x11\x04\xd9\x43\x81\x16\x3d\xd0\x22\x25\x11\x2b\x93\x2c\x49\xe5\x81\x20\xff\xbd\x33\x24\xf5\x88\x63\x6f\x60\x77\xd7\x0b\x7b\xa5\xe1\xc7\x99\x6f\x1e\xfc\x98\xdc\xba\xe7\x96\x17\x33\xc7\x88\x36\x9c\xbc\xcc\x08\x7c\x98\xb0\xba\xa5\xcf\x6b\x22\x64\x2b\x24\xbf\xf1\xc6\x2d\x2d\xbf\xd5\x46\x75\x92\xad\x89\x54\x83\x55\x19\xc6\xcd\xd4\xa2\x29\x63\x42\xd6\x6b\x72\x1d\xde\x4b\xd5\x2a\x00\x5c\x5e\x5f\x47\xc3\x63\x23\x1c\x5f\x59\x4d\x4b\xbe\xc6\xa0\xab\x47\x43\xf5\xcd\xec\x75\xe6\x9a\x18\x7f\x1a\xea\xb2\xfa\x84\xff\x70\x3d\x95\x8a\xf1\x95\xea\x9c\xee\x5c\x44\xee\xa8\xa9\x85\x5c\x39\xa5\xd7\xe4\x93\x7e\x42\xd4\x25\xa2\x4c\x27\x2d\x71\x66\xdd\xa8\x07\x6e\x00\x7b\x51\x76\xc6\x22\x0d\xad\x84\x74\xdc\x20\x30\xcf\x62\xee\xb9\x2d\x8d\xd0\xae\x98\x5d\x2d\xaa\x4e\x96\x4e\x28\xb9\x58\xc6\x00\x57\x8b\xe4\x6f\x46\x1d\x85\x10\x75\xdd\xf2\xcd\xdc\x29\xd5\x3a\xa1\xe7\xff\x24\xcb\x34\x3e\x2f\x96\x37\x11\x3b\x9f\x06\x9f\x2f\xd3\xb2\x15\xe5\xb7\xd1\x29\xef\xbd\xfa\x32\x08\xc9\xd4\x63\xda\xaa\x92\xe2\x62\xda\x18\x5e\x91\x0d\x38\xe1\xa9\x83\xac\xb8\x5b\xa6\x9a\x1a\x2e\x9d\x5d\xcc\xbd\xb3\x0a\x36\x2c\x12\x68\x14\x85\xd0\xd4\x39\xb3\x98\xe3\x9e\xf9\xd2\xbb\x7c\x05\x12\xf8\x85\xac\x62\x36\x39\x13\x0f\xa4\x6c\xa9\xb5\x9b\xa4\x54\xd2\x51\x68\xa5\x49\x0a\x40\xe7\x95\x32\x3b\xb2\xe3\xae\x51\x6c\x93\x68\x65\x9d\x37\xc3\x82\xa3\xdb\x96\xf7\x9b\xc2\x8b\xff\x5d\x81\x03\xc6\xa5\xe5\x2c\x22\x11\x6b\x8a\xd9\x45\xee\x9a\xe2\xb3\xda\xed\xa8\x64\x79\x06\xcf\x68\x61\x45\x0e\x6d\x2d\x5e\x5e\x48\x7a\x0b\xe5\x48\xe3\x3a\x79\x7d\xcd\x33\x5c\x00\x20\x1b\xbc\x64\xe8\x66\xdf\xe3\x57\xc7\xa0\xcd\x47\x1d\x86\xe5\xd3\xfc\x71\x63\xbe\xe7\x0f\x96\x4f\xf2\x47\x5d\x67\x27\xfe\x66\x17\x84\x80\x3b\x51\x11\xfe\xef\xe0\x14\x31\x24\xf9\x0a\xc3\xa9\xa1\x6e\xe0\x1e\x51\xf9\xb6\x73\x4e\x49\x82\x2d\xa0\x7e\x2c\x36\x49\x86\x53\x93\x0d\x74\x6e\xe9\x8e\x03\x1a\xc6\x93\x1a\x97\xf4\xdd\xd8\x3a\x49\xe0\xbb\x7a\xb2\xfe\x3f\xdb\x95\x25\xb7\x36\x41\x2e\x06\x4a\x15\xdc\xf6\x44\x78\x6b\xf9\x39\x01\x95\x3e\x16\x8f\x51\x59\xe3\xf8\x60\x3a\xfb\xd1\x0e\xa5\x7d\x47\x3b\x1b\xb3\x26\xc3\xe7\x04\x2e\x86\xdb\x6e\xc7\x3f\xcc\xfe\xde\xc3\x46\x42\x63\xac\x49\x15\xce\x22\xa0\x31\x81\x8f\xaa\xe1\xb3\x7c\x57\x7c\x3f\xeb\x7b\x2f\x1f\x0f\xd5\xaf\xc0\xe7\x81\x13\x64\x73\x70\xb2\x02\xbd\x88\x8a\xbd\x9d\x1c\x71\x2f\x8e\x20\x3d\xc9\xb4\x08\x00\xa1\x04\x55\xe2\x78\x9d\x3b\x39\x1a\x83\xf3\xf4\xcb\x6f\xb0\x92\x14\x97\x07\xed\x79\x46\x0b\xb2\xb7\x02\xe6\x5f\xe4\xd6\xea\x9b\xe9\x2f\x12\xf4\xf1\x8f\x94\x90\x57\xb4\x6b\x61\xbe\xcf\xa3\x07\xc3\x8a\x22\x90\x14\x39\xdc\x24\xb2\x8f\x51\xb7\xcf\xba\x11\x20\x56\x64\x78\x5a\x55\xa2\xe5\x00\xcb\x10\x57\x10\x1b\xa5\x85\x16\x3f\x9f\x1f\x88\xca\x39\xfc\xbc\x54\x05\x7e\x79\x06\x1d\x3e\x74\xa8\x45\x71\x0b\x37\x6e\x9e\x89\xfd\xa1\x3b\x36\x67\x00\xb9\xc2\x14\xc8\x7a\x13\xd8\xf6\x07\x03\x40\xa8\xf0\xc5\xec\x83\x1b\x20\xdc\xf4\x9c\xc5\x57\x7f\xb3\x26\x44\xb0\x30\x79\x78\xe7\x1d\xb8\x1a\xee\x3b\x39\xcc\x72\x53\xdc\x09\x36\x79\xfb\xfd\x49\x38\x48\xf7\x8d\x92\x36\x41\xce\x38\x7b\x63\xf1\x02\x3a\xb1\xfc\xa1\xea\xb8\xe5\x70\x9e\x98\x9e\x6f\x50\x9f\x6b\xdf\xae\x11\x63\xf0\x00\xc7\xc5\x7b\xbc\xaf\x87\x45\x20\xdf\xa7\x1f\xce\x9d\x54\x8e\xa4\x91\x56\xfa\xc5\xfe\xc5\x8d\x02\x74\x3c\x93\x91\xdc\x68\x17\xb2\x52\x63\xaf\x02\xaa\x06\x07\x7f\x52\xe1\x82\x36\xa6\x98\x78\x94\xc9\x6b\xc0\x04\x2d\x19\xf7\x44\x6d\x1b\x7a\xfa\xfe\x21\xe9\x6f\xb1\x77\x67\x7b\x4c\x7d\x32\x9b\xd3\xe3\x3c\x1c\xe1\x30\x27\xde\x0d\xda\x3f\xef\x58\x7a\x67\x14\x06\x4e\xa1\x4d\x1e\x34\x05\x8c\x75\xd8\xcb\x17\x37\x1f\xce\x2d\x10\x0e\x39\x4d\xe6\x35\x64\x71\xdc\xff\xfb\x3a\xf7\xc6\x69\x09\x4e\x60\x17\x8d\x47\x76\xfb\xd3\x74\xaa\x06\x1c\x2f\xf3\x8f\x91\xa5\x1f\x4c\xe8\xff\xea\xd0\x77\x34\x25\xca\xce\x1b\x21\xc1\x47\xbc\x60\xe1\x4f\xd0\x20\x5f\xff\x05\x00\x00\xff\xff\x7f\x8f\x51\x5a\x65\x0c\x00\x00")

func assetsTemplatesNodeHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesNodeHtml,
		"assets/templates/node.html",
	)
}

func assetsTemplatesNodeHtml() (*asset, error) {
	bytes, err := assetsTemplatesNodeHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/node.html", size: 3173, mode: os.FileMode(420), modTime: time.Unix(1456762016, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesNotfoundHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x54\x41\x6f\xdb\x3c\x0c\x3d\x7f\xfd\x15\xaa\xbe\xeb\x64\x21\xdd\x65\x07\x3b\xc0\x36\xec\xb0\xdb\x0e\x1b\xb0\x2b\x6d\xd1\xb6\x32\x59\x52\x25\x3a\x6d\xfe\xfd\xa8\xd8\x71\xd3\xa2\x03\x8a\x61\x87\xc0\xa2\x48\x3e\x92\x8f\x4f\xa9\x6f\x4d\xe8\xe8\x14\x51\x8c\x34\xb9\xfd\x4d\x5d\x3e\xc2\x81\x1f\x1a\x89\x5e\xee\x6f\x84\xa8\x47\x04\x53\x0e\x7c\x9c\x90\x40\x74\x23\xa4\x8c\xd4\xc8\x99\x7a\xf5\x41\x5e\xbb\x46\xa2\xa8\xf0\x7e\xb6\xc7\x46\xfe\x54\x3f\x3e\xaa\xcf\x61\x8a\x40\xb6\x75\x28\x45\x17\x3c\xa1\xe7\xbc\xaf\x5f\x1a\x34\x03\x3e\xcb\xf4\x30\x61\x23\x8f\x16\x1f\x62\x48\x74\x15\xfc\x60\x0d\x8d\x8d\xc1\xa3\xed\x50\x9d\x8d\x77\xc2\x7a\x4b\x16\x9c\xca\x1d\x38\x6c\x76\x0c\xb4\x20\x91\x25\x87\x7b\x83\x53\x10\x98\x52\x48\xb5\x5e\x6e\x56\xb7\xb3\xfe\x97\x48\xe8\x1a\x99\xe9\xe4\x30\x8f\x88\x5c\x69\x4c\xd8\x37\x52\xeb\xce\xf8\x43\xae\x3a\x17\x66\xd3\x3b\x48\x58\x75\x61\xd2\x70\x80\x47\xed\x6c\x9b\x35\x3d\x58\x22\x4c\xaa\x0d\x81\x32\x25\x88\xfa\x7d\xb5\xab\x76\xba\xcb\x59\x6f\x77\x15\x5b\x5b\x37\xb9\x4b\x36\x92\xc8\xa9\x7b\x03\xfc\xe1\x7e\xc6\x74\xd2\x77\x67\xcc\xc5\xa8\x26\xeb\xab\x03\xe3\xd5\x7a\x81\xda\xff\x05\xee\x9f\xda\x3e\x5c\x77\xfd\xbc\xc8\x1b\xc8\x2a\x43\x1b\xec\x61\x76\xb4\x8e\xcc\x19\xfa\x22\x94\xba\x0d\xe6\xb4\x36\xeb\xe1\x28\x3a\x07\x39\x37\x92\x8f\x2d\x24\xb1\x7c\xd4\x9a\x7e\x31\x7b\xfb\x88\x46\x51\x88\x52\xa4\xc0\x4b\x2d\xd1\x76\x60\xe9\x04\xbf\xea\x84\xc1\x8c\xdd\xc0\x8a\x3e\xc0\x7a\x9e\xac\x77\xb3\x35\x1c\xf3\x5f\x7d\xab\x94\xf8\x94\xc0\x1b\x51\x7e\x14\x86\xc1\xa1\x18\x90\xc4\x90\xc2\x1c\xd1\x88\x3e\x24\xd1\x62\xe1\x43\x4c\xa1\xb5\xec\x35\x36\x47\x07\x27\xa1\x54\x01\xb8\xc2\x5f\xdb\x2a\x23\x61\x2a\xe8\x65\xac\x99\x28\x78\x51\x9e\x4b\x23\x17\x43\xbe\x88\x5f\x8a\x4a\x61\x80\x60\x35\x4a\xaf\xce\x41\xcc\xdb\x35\xa4\xa1\x3c\x9f\xff\xdb\xac\xf0\x11\xa6\xe8\x50\xad\xe9\x97\x48\xb5\x5b\x4a\x96\x6d\x47\xf0\x97\x22\x39\xa9\xe0\xdd\x49\xee\xbf\x2f\xb3\x3d\x71\xc4\xdb\xe3\xb8\xd7\x72\x2c\x33\xa5\x18\xfb\xbc\xe1\x7f\x1d\x53\xeb\x85\x86\xc5\x80\x17\x64\xb4\x65\x17\x9b\x66\xe4\xf9\x61\xd6\x1a\x0a\xd3\x9a\xa9\xde\xd6\xfa\x64\xd4\x9a\x53\x2f\x0a\x7c\x6d\xdb\x4f\x5a\x18\xef\xf6\xdf\x60\x60\x0e\x02\xf1\x5e\x67\x6f\x58\x7f\x77\x17\x94\x15\x90\xdb\x3b\x2b\x91\x5d\xe7\x7f\xb7\xdf\x01\x00\x00\xff\xff\x11\xbe\xda\xcb\xee\x04\x00\x00")

func assetsTemplatesNotfoundHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesNotfoundHtml,
		"assets/templates/notfound.html",
	)
}

func assetsTemplatesNotfoundHtml() (*asset, error) {
	bytes, err := assetsTemplatesNotfoundHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/notfound.html", size: 1262, mode: os.FileMode(420), modTime: time.Unix(1456714925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _assetsTemplatesRunHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x55\x4f\x6f\x9c\x3e\x10\x3d\x67\x3f\x85\x45\x22\x65\x73\x00\xf2\x8b\xf4\xbb\x6c\xbd\x1c\xfa\xe7\x10\xa9\x8a\xa2\xe6\x50\xa9\x55\x0f\x5e\x3c\x80\x55\x63\x23\x7b\x48\x77\x15\xe5\xbb\x77\x6c\xd8\x0d\x25\x55\xb3\x6d\xa5\x6a\x05\x3b\xcc\x3c\xbf\x61\x1e\xf8\xc1\x3d\xee\x34\x14\x0b\xc6\x50\xb2\xce\x01\x7b\xa0\x50\x2a\xdf\x69\xb1\x5b\x31\x65\xb4\x32\xf0\x8a\x52\x1b\x51\x7e\xad\x9d\xed\x8d\x5c\x31\x63\xc7\x9c\x75\x12\xdc\xd3\x75\x27\xa4\x54\xa6\x5e\xb1\xcb\x70\xf5\x48\x47\x86\x62\xa3\x81\x61\x13\x69\xa7\x1c\xa7\xd5\xff\xe1\x77\x00\xfa\xd2\x59\xad\xc1\x45\x60\x2b\xb6\x69\x03\xaa\x6e\x70\xc5\xfe\xbb\xba\xec\xb6\x01\x66\xef\xc1\x55\xda\x7e\x4b\xe9\xbe\x06\xf4\xb0\x98\xe7\xe3\x08\x9c\xb2\xaa\xc3\x30\xcb\xd9\xb2\xea\x4d\x89\xca\x9a\xe5\x45\x64\x3c\x5b\x26\x9f\xa5\x40\x91\xa2\xad\x6b\x0d\xeb\x73\xb4\x56\xa3\xea\xce\xbf\x24\x17\xd9\x18\x2f\x2f\x22\x21\x9d\x89\x72\xa4\xe2\x52\xdd\xb3\x52\x0b\xef\xd7\x49\x69\x0d\x0a\x92\xc3\x25\xa1\x05\x6f\xae\xf6\x85\x87\x07\xa6\x2a\x52\x01\x59\x76\x63\x25\x7c\xe8\x4d\x76\x87\xc2\x21\xc8\xec\xda\x7f\x02\x67\xd9\xe3\xe3\x80\x99\xd4\x6d\xd7\x4d\xeb\x08\x5b\x4c\x95\xa9\x2c\x01\x41\x7b\x38\x2c\xa9\x27\xac\x1f\x85\x42\x62\xc6\xde\x67\xef\xb6\xfb\x90\x5d\xee\x97\x4b\x61\x6a\x70\x4f\x04\x31\xe9\xfb\xb2\x04\xef\x43\xd6\xc8\x81\x75\x16\x24\x05\x85\xb1\x47\x76\x23\xda\xb0\x90\x9d\xee\x33\xa1\xeb\xf5\x5b\x4a\xf1\xbc\xb9\x8a\x63\x57\xd6\xb5\xac\x05\x6c\xac\x5c\x27\x9d\xf5\x18\xd5\xa0\xc2\xf0\xa8\x47\x49\xc6\xe7\x1e\xce\x29\xe9\x26\xc1\x78\x90\x23\x32\x60\x5d\xb1\x38\xe1\xd8\x14\x6f\x6c\xdb\x0a\x23\x79\x4e\x71\xc8\xc8\x82\xd3\x4b\x58\x4c\xdb\x8f\x90\x78\x0f\xa1\x46\x58\x79\x20\xca\x03\xd3\x9c\xf4\x0e\xa5\xed\x71\xc2\xb9\x38\xa1\xf2\x9c\x77\x40\x1d\x68\x59\xca\x9e\x57\x5f\xf7\x55\xf6\x1e\x4c\x90\x64\xb3\x43\xf0\x8c\x8b\xfd\x84\x1b\x34\x8c\x8e\x74\xeb\xe3\x9f\x84\x4a\xf4\x1a\x13\xd6\x38\xa8\xd6\x49\x6e\x88\x27\x9f\xeb\x9a\xbb\xde\xe4\xcf\xa4\xa5\xf7\x37\xf4\x4a\x0a\xee\x3b\x61\xf6\xfc\xb5\xde\x75\x8d\x22\xed\xd8\x21\x4a\x2b\xa5\x81\x60\x79\xc0\x15\xcc\x8f\x63\x8a\x30\xe5\x31\xa2\x80\x73\x47\x88\x42\xa8\x5f\x88\x42\xd5\x7f\x26\x0a\xf5\xfa\x13\x51\xe2\x98\xc7\x8a\x12\xb7\xe9\x44\x95\x23\xf7\xf2\xbc\x38\xdd\x4f\xc7\xb4\x8d\xbb\xff\xa5\xb6\x33\x8b\xf8\xb1\x6d\x2c\xfe\x5e\xdb\x5b\x35\x6b\xf9\xb4\xc7\x5a\x99\xdd\x3a\x1b\x8c\x22\x23\xd4\x51\x6c\xc1\x81\x48\xee\x60\x41\x7f\x31\xc8\xcf\x2d\x6d\x18\x6b\x70\x31\xae\x8a\x1b\xfa\xc0\xf0\x5c\x15\x2f\xcc\x4a\x51\x30\x9c\x68\x53\x79\xf0\x29\xf2\xef\x9c\x0c\xbc\x58\x7c\x0f\x00\x00\xff\xff\x2b\xc4\xd3\x55\xe6\x06\x00\x00")

func assetsTemplatesRunHtmlBytes() ([]byte, error) {
	return bindataRead(
		_assetsTemplatesRunHtml,
		"assets/templates/run.html",
	)
}

func assetsTemplatesRunHtml() (*asset, error) {
	bytes, err := assetsTemplatesRunHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "assets/templates/run.html", size: 1766, mode: os.FileMode(420), modTime: time.Unix(1456762014, 0)}
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
	"assets/css/default.css": assetsCssDefaultCss,
	"assets/templates/cluster.html": assetsTemplatesClusterHtml,
	"assets/templates/error.html": assetsTemplatesErrorHtml,
	"assets/templates/layout.html": assetsTemplatesLayoutHtml,
	"assets/templates/log.html": assetsTemplatesLogHtml,
	"assets/templates/node.html": assetsTemplatesNodeHtml,
	"assets/templates/notfound.html": assetsTemplatesNotfoundHtml,
	"assets/templates/run.html": assetsTemplatesRunHtml,
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
		"css": &bintree{nil, map[string]*bintree{
			"default.css": &bintree{assetsCssDefaultCss, map[string]*bintree{}},
		}},
		"templates": &bintree{nil, map[string]*bintree{
			"cluster.html": &bintree{assetsTemplatesClusterHtml, map[string]*bintree{}},
			"error.html": &bintree{assetsTemplatesErrorHtml, map[string]*bintree{}},
			"layout.html": &bintree{assetsTemplatesLayoutHtml, map[string]*bintree{}},
			"log.html": &bintree{assetsTemplatesLogHtml, map[string]*bintree{}},
			"node.html": &bintree{assetsTemplatesNodeHtml, map[string]*bintree{}},
			"notfound.html": &bintree{assetsTemplatesNotfoundHtml, map[string]*bintree{}},
			"run.html": &bintree{assetsTemplatesRunHtml, map[string]*bintree{}},
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
