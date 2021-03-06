// Code generated by go-bindata.
// sources:
// data/aws-simple/build/build-php.sh
// data/aws-simple/build/template.json.tpl
// data/aws-simple/deploy/main.tf.tpl
// data/common/dev/Vagrantfile.tpl
// DO NOT EDIT!

package phpapp

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path/filepath"
)

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	name string
	size int64
	mode os.FileMode
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

var _dataAwsSimpleBuildBuildPhpSh = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x7c\x55\xdd\x6e\x1b\x37\x13\xbd\xdf\xa7\x98\x28\xc6\x97\xaf\x45\x28\x2a\x06\x5c\x14\xf2\x0f\xaa\xc8\x4a\x6d\x20\x90\x0c\xa9\x4d\x2f\x5c\x23\xa0\x96\xb3\x2b\xc6\xbb\x24\x43\x72\x25\xcb\x8e\xde\xbd\x43\xae\x2c\x5b\x2e\xea\x2b\x71\xc9\xf9\x39\x73\xe6\xcc\xe8\xed\x1b\x3e\x57\x9a\xcf\x85\x5f\x64\x99\xc7\x00\xcc\x80\x36\x8d\xde\x1e\xd1\x39\xbc\x53\xe9\x68\x95\xc5\x42\xa8\x6a\x7b\x1d\x9c\xc8\x31\xcb\xe8\x64\xdc\xff\x7f\x82\x87\x0c\x00\x2a\x93\x8b\x0a\xbc\x69\x5c\x8e\x85\xaa\xf0\xf4\xe0\xc3\xd3\x75\xa5\x34\x6a\x73\x7a\x70\x18\xaf\x30\x5f\x18\xe8\x8c\xa6\xd3\xc9\x14\x44\x80\x83\x87\x27\xa7\x4d\xff\xe0\xa1\xb5\xdd\x1c\xc3\x67\xe1\x03\xf9\x97\xbe\xdf\x89\x6e\xa5\x43\x0b\x26\x04\x03\x7c\x29\x1c\xa7\x07\xee\xd7\x9e\x7e\xe0\x07\x84\x84\x4d\xc3\x61\x2f\xdb\x64\x84\xce\xc2\xbb\x04\x0e\x3a\x07\x0f\x1f\x07\xb3\x8b\xaf\xb3\xc9\x9f\xd3\xe1\x68\xd3\x89\x17\x9f\x2f\xc7\xa3\xf1\x64\xd3\x79\x07\x84\x21\xcb\x0c\xc6\x12\xe8\xe1\xb7\x0e\x1c\x9e\xfd\xef\x03\x85\xa3\xa0\x25\x3a\x60\xa1\xcd\x77\x06\x5c\xe2\x92\xeb\xa6\xaa\x8e\x61\x93\x99\x2a\x39\xb4\x65\x5c\x47\x8b\x1b\x20\xe7\xf8\x94\xbd\x85\xbc\x32\x8d\x64\xb9\xd1\x85\x2a\x21\x17\x1a\x94\x0e\xe8\x0a\x74\x08\x2b\x15\x16\x20\x6c\x80\xdc\xd4\xb5\xd0\xd2\x83\x2a\x40\x85\x77\x1e\x7c\x50\x55\x45\x96\x60\x9d\xa1\x3a\xbd\xa7\x24\xd0\xf9\x4b\xa8\xa0\x74\x09\x05\x15\xb2\x17\x96\x30\x51\x08\x5b\x61\xc0\x6e\xb7\xdb\xc9\x1a\x4d\xfe\x70\x7d\x0d\xac\xd8\x92\xa3\xe6\x3c\x79\x70\xa5\x7d\x10\x3a\x47\x3e\x37\x26\xb0\x42\x69\xe5\x17\x28\xe1\xe6\xe6\x18\xa4\x21\x5a\x7d\x85\x44\x6b\xaf\x7b\x94\x49\xa3\xa9\xa7\x31\xef\x40\xca\x98\x36\x22\x25\xce\x8d\x57\xc1\x38\x85\x1e\x08\x32\x34\x56\x8a\x08\x2a\xe5\xc5\x3b\x6b\x5c\x80\xf3\xd1\xc7\xcb\xc1\xf8\xeb\xa7\xe9\x64\xfc\xc7\x68\x7c\x7e\xaa\x8d\x4e\x45\x8b\x3c\xa8\x25\x12\xc1\xe0\x1b\x69\x62\x3c\x56\x92\xb4\x52\x08\x04\xb6\xfe\xd7\x4b\x02\x4b\x44\xb0\x35\x09\xa9\x08\x2b\xe1\x90\x11\x23\x16\x5d\xa0\xfc\x2c\xd2\x66\xf4\x93\x97\x94\x2c\x7a\xee\x30\xae\xa3\xa3\xb5\xa2\x6f\xb4\x74\xf8\x8d\xdb\x85\x3d\x62\x47\xdd\x5f\xa8\x2b\x33\xc4\xda\x47\xde\xe6\x48\x35\x7d\x6f\x94\x23\x12\xe8\xd3\x3a\x5c\xa2\x0e\xd0\x69\xb4\x68\xc2\x82\x8e\x2a\x27\x70\x12\xac\xc8\x6f\x45\x89\xbe\x43\xce\x49\x4a\x1e\x4c\x43\x7a\x28\x5e\x82\xed\xee\x55\x71\x8b\xeb\x6d\x7d\xaf\x94\x9d\x48\xbe\x6c\xdd\x23\xd1\x57\x17\x57\xf0\x65\x38\xf3\xef\x61\x40\x69\x17\xf8\x3e\x31\x6d\x08\x8e\xdb\xe1\x48\x7c\xbf\xc2\x97\x48\x9e\x87\x10\x6b\xa6\x71\x9b\x6f\xbf\x59\x6d\x24\x4b\x77\x7f\x53\xb3\xe7\xf7\x0e\x4a\x9a\xe8\x1a\x5d\xde\x38\x45\x83\x39\x6f\x54\x25\x19\x09\x2e\x56\x4e\xdf\xd1\x8a\x9e\xda\x43\xe2\xaf\xce\xdd\x9a\x84\xd0\x9e\xd7\xfe\x7b\xd5\x1e\x0b\x5b\xb7\x87\x52\xb6\xbf\x0e\x85\x8c\xa3\xdb\x7e\xd9\x92\x2c\xdb\x42\x47\x77\x21\x29\x21\x29\xca\xa6\x3a\x52\x11\xf5\xad\x54\x34\x63\x16\xb8\x77\x4b\x1e\xc7\x88\xba\x69\xdb\xb7\x20\x1c\xdc\xdf\x91\x98\x43\x6d\x77\x4f\xdd\x50\xde\x03\x1b\xbe\xb0\x4f\x39\x66\x18\x52\x02\x52\x4a\xad\xbc\x57\x46\xef\x13\x46\xa3\xba\xd2\xc0\xa6\xb0\x5a\xad\x18\xb5\x41\xf4\x5f\x46\xa1\x49\x6c\xc7\xe7\xf9\x35\x8f\x53\x66\x3c\xba\xee\x37\x6f\x34\xd0\xcc\x44\x89\x10\x33\x2f\x1a\x38\x7c\x34\x8b\x39\x89\x40\x99\x80\x3f\x52\xc9\xfc\x0c\x16\x21\x58\xdf\xe7\x9c\xba\xb6\x8b\x69\x5c\xc9\xb7\x1d\xa4\x46\xff\x88\xc4\xc5\xd0\x5b\xc8\xf5\x12\x76\x96\x76\x41\x7c\xf0\xc6\xc7\xad\x47\xeb\x34\x6d\xed\xc7\xc7\x6c\x0b\x67\xda\x68\x1d\xb1\x3c\xde\xef\xf4\xd9\x62\x7a\x0c\xcb\x9a\x1d\x07\xff\x15\xf1\x49\x56\x6c\x65\xdc\x2d\x05\x65\xb1\x53\x7b\xcc\xd0\x9b\x36\x8c\xb6\x62\x7b\xd8\x8d\x3b\x8d\x67\xa1\xb2\xb6\x29\xc3\xb4\xac\x48\x67\xa9\xf3\x51\x8e\x7b\x3d\x71\x35\x70\x62\x83\x6f\x95\xca\x69\x84\x69\xc6\x51\x8b\x79\x85\x92\xf7\x7a\x3d\x0a\x5f\x88\xa6\x0a\xdd\xb8\xf5\xb2\x8c\xe6\x12\x4e\x4e\x06\x57\x83\xe1\xc5\x68\x38\x19\x7f\x22\xc6\x5a\xad\x20\xbe\x16\x68\xa7\x9e\x18\xe5\xf9\x2a\xa7\xc1\xfe\x1d\x35\xc1\x8e\xe3\x3e\x5f\xc3\x84\x0c\xb3\x93\x2f\xca\x85\x46\x54\x17\x86\xfe\x7b\x7e\xee\xff\xda\x3b\x23\xee\xce\x4d\xde\xd4\x34\x20\x53\x5a\xa1\x2f\x75\x03\x70\x72\x4e\xcb\x24\x4f\xeb\x67\x5f\x3c\xd1\x15\x60\xda\x6e\x1b\x88\x84\x96\x4e\x10\x51\x32\x3a\xf1\x9d\xd7\x59\x76\xc2\x9f\x65\x3d\xcb\x9e\x4a\xdc\x12\x49\xb4\xc5\xe5\xfc\xa6\x93\xfd\x13\x00\x00\xff\xff\x89\xa6\x02\xdd\xb2\x07\x00\x00"

func dataAwsSimpleBuildBuildPhpShBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildBuildPhpSh,
		"data/aws-simple/build/build-php.sh",
	)
}

func dataAwsSimpleBuildBuildPhpSh() (*asset, error) {
	bytes, err := dataAwsSimpleBuildBuildPhpShBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/build-php.sh", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleBuildTemplateJsonTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\x54\x5d\x6f\x9b\x30\x14\x7d\xe7\x57\x58\x48\xe9\x53\x80\x6c\xad\xa6\x69\xaf\xfb\x19\x55\x44\x0d\xb8\xe1\x2a\xb6\xb1\x7c\x4d\xa6\x16\xf9\xbf\xef\x9a\xef\x4c\x0d\xcd\x9a\x17\x23\x9f\xc3\x39\xc7\x87\xeb\x74\x11\xa3\x5f\xac\x40\xe7\x86\x97\x67\x61\xf3\x8b\xb0\x08\x8d\x8e\x7f\xb1\xf8\x90\xfe\x4c\x0f\xf1\x3e\x1a\x38\x17\x6e\x81\x17\x52\x20\x41\xc3\x6b\xb4\xc9\xff\x60\xce\xcb\x52\x20\xe6\x67\xf1\x46\x88\x6e\xa5\xdc\xaf\x51\x14\xa5\x15\xee\x16\x6a\xc5\x69\x30\xbb\x42\x50\xb6\x27\xca\xe3\xea\x11\xe8\xf7\xfd\x14\xc4\xd8\xe6\x02\x21\x23\x25\x25\xc2\xf3\xf8\x56\xb7\x63\xaf\x8d\x65\x15\x58\x06\x9a\x1e\x5b\x5d\x71\x47\xac\x9c\x76\x30\x2d\x5a\x90\x15\xdb\xf9\x89\x3c\xae\x24\xe7\xde\x8c\x08\xa7\xc5\x5a\x48\x19\xef\x17\x00\xb4\x04\x1d\xa0\xe7\x58\x9d\x83\x6c\x62\x58\xe6\x94\xc9\x1a\xe7\x9a\x6c\x31\x48\xba\x2e\x38\xcb\xa6\x31\xe9\x6f\xda\x75\xc2\x32\xef\xe3\xe3\xa8\xe4\xf7\xb7\x3d\x5f\x41\x8a\xb5\x25\x36\xad\x2d\x7b\x84\x34\x83\xa5\xf7\xd9\x1a\xaf\x04\x3a\xd0\xbd\x6b\x20\xfd\x47\x9a\x3b\xc2\x6c\x15\x50\x56\xf7\x1e\xdd\x7b\xf6\xf0\xc0\x0a\x8e\x35\x4b\x33\xc5\x41\xa7\x58\x7f\xd0\xc5\x8e\x09\x5d\x85\xef\xb5\xf5\x49\x36\xea\xd9\x31\x1a\xd4\x82\x42\x28\x52\xa0\x14\x2d\xd2\x39\x5f\xe6\xc1\x79\xa1\x33\x0f\x1e\x2b\xda\x3d\x4d\x26\xdc\x98\xd4\x9d\xde\xbf\x54\x18\x96\x16\x8c\x0b\x50\x3f\x6e\x89\xa9\x4d\x38\xfd\x24\xd5\xaf\xc7\x69\x8c\x7b\xca\x38\xc2\xf3\x7d\xd2\x5c\xf5\xd2\x21\xca\xac\x3c\x1b\x72\xc5\xdf\xa9\x74\x51\xe0\x82\x5d\xdd\xbe\x5b\xbd\x5c\x5f\xd3\xed\x72\xe2\xab\x1b\xbb\xa5\xb8\x10\x3f\x51\x9c\x6f\xf9\x96\xda\x40\xfa\x2c\x5b\x3f\x01\x39\x57\x30\xf4\x01\xc9\xf7\x6f\x3f\x1e\x0f\xd5\xd3\xd3\xc2\x01\x8d\x8e\x6b\x62\x4d\xb5\x95\x8f\xa9\xe4\xf6\x24\x56\x32\x58\xe7\xc1\x79\xaa\xbb\x2d\x68\x76\xdb\x55\xa9\x0a\xf2\x09\xeb\xba\xf0\x44\x63\xfd\x6f\x76\x5a\x69\x88\xb8\x32\x1f\x25\x1e\xfe\xb2\x8e\x51\xe4\xa3\xbf\x01\x00\x00\xff\xff\x4f\xd6\x14\x86\x64\x05\x00\x00"

func dataAwsSimpleBuildTemplateJsonTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleBuildTemplateJsonTpl,
		"data/aws-simple/build/template.json.tpl",
	)
}

func dataAwsSimpleBuildTemplateJsonTpl() (*asset, error) {
	bytes, err := dataAwsSimpleBuildTemplateJsonTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/build/template.json.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataAwsSimpleDeployMainTfTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x52\xc1\x8e\xdb\x20\x10\xbd\xe7\x2b\x46\x74\x8f\x5d\x27\xed\xb1\x52\xcf\xbd\xb5\x1f\x50\xad\x10\xc6\x24\x45\x8b\x01\xc1\x90\xca\xb2\xfc\xef\x1d\xa0\x16\xc1\xbb\xbd\x36\x39\xf9\xcd\x63\xde\xcc\x9b\xf7\x01\xbe\x29\xab\x82\x40\x35\xc1\xb8\xc0\x0f\x44\xf7\x11\x26\x07\xd6\x21\xa8\x49\x23\xcc\xc2\x26\x61\xcc\x72\x3a\xdd\x45\xd0\x62\x34\x0a\x98\xb6\xd7\x20\xb8\x9e\x18\xac\xdb\x03\x2c\x7e\x47\x2e\xa4\x54\x31\xf2\x57\xb5\xbc\x53\x8c\x4a\x06\x85\xff\x28\x06\x75\xd3\xce\x1e\x0a\x44\xe5\x56\xcc\xaa\xc0\x8f\x0f\x66\x7d\x60\x6a\x1b\x51\x58\xa9\x38\x2e\x3e\xd3\x61\x52\x57\x91\x0c\xc2\x57\x60\xf8\x79\x98\xb5\x0c\x8e\xc1\xe3\x8b\x98\x46\x4b\xd3\xf8\x34\x1a\x2d\x0f\xdd\xee\x5e\x72\xa9\xa7\xf0\x0e\xfc\x77\xed\x93\x0f\xee\xae\x27\x15\xca\xf4\x04\x9d\x00\xda\xf2\x59\xf5\x69\xa5\x87\x43\x6f\xca\xc6\x88\xd6\x6c\xe8\x69\x0d\x2f\xb4\x6a\x08\xe4\x5f\x47\xab\x38\x51\x68\x88\xa0\xa2\x4b\x41\x36\x7f\x53\xd0\xb8\xf0\x5b\x70\xc9\x33\x02\xbd\xaf\x93\x65\x0f\x6b\x9f\x75\xad\x1f\xdb\xf6\x5c\x5b\xee\xc7\x2c\x9a\x75\xc1\xa6\x57\xbf\xa9\x44\x35\x6d\x6f\x24\x17\x4b\x3f\x00\x5a\x1f\x9d\x74\xa6\x8e\xf7\xfc\xa9\x80\xd7\xe0\x66\xee\x5d\xc0\x02\x5e\x0a\x86\x6e\x47\x1a\x96\xad\xe5\xa3\x71\xf2\x35\x12\xf6\x93\x5d\x86\xf2\x3f\x5f\xd8\x0b\xd5\xb7\xac\xa6\xfe\x9b\xd8\x1b\x1b\xf7\x28\x3d\x1a\x48\x81\x83\xf6\x6b\xf7\x98\x75\xf1\xad\x4b\x5f\x2b\x77\x70\xbd\x7d\x0d\x1d\x79\xdc\xf5\xe9\xb2\x58\x88\x7b\xf2\x0f\x82\x3b\x5c\x4f\x92\xcf\xd3\x1f\x9d\x3a\xd7\x2d\x9f\xd6\xb7\x89\x18\x68\x9d\x21\x9f\xf3\x25\x3f\x46\x71\x23\x7f\xe1\x7b\x16\xe9\x82\xc1\xaa\x29\x2e\xa1\x4f\x08\x2c\x05\x53\x3d\xb8\x0b\x93\x0a\xf5\x17\xa2\xff\x72\x3e\x57\x89\x7d\xc7\xd2\xbc\x2e\xc0\x27\x1b\xb7\x73\x0e\xe8\x9f\x00\x00\x00\xff\xff\x5f\x73\x79\x4b\x5f\x04\x00\x00"

func dataAwsSimpleDeployMainTfTplBytes() ([]byte, error) {
	return bindataRead(
		_dataAwsSimpleDeployMainTfTpl,
		"data/aws-simple/deploy/main.tf.tpl",
	)
}

func dataAwsSimpleDeployMainTfTpl() (*asset, error) {
	bytes, err := dataAwsSimpleDeployMainTfTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/aws-simple/deploy/main.tf.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _dataCommonDevVagrantfileTpl = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x55\xdb\x6e\xdc\x36\x10\x7d\xd7\x57\x4c\x64\xe7\x06\x58\x52\x63\x34\x79\x70\x62\xa3\x8e\xe3\xd4\x06\x52\x3b\xf0\x3a\x7d\x69\x8b\x0d\x57\xe2\x4a\xac\x25\x92\x21\xa9\x4d\xb6\xeb\xfd\xf7\x9e\xa1\xb4\xbe\x24\x41\x5a\xc0\xc6\x52\xc3\xb9\x9e\x99\x39\xdc\xa2\x5f\xa5\x96\x4e\x04\x59\xd1\x6c\x49\xe7\x21\x98\x1d\xaa\x0c\x69\x13\x48\x56\x2a\x3c\x48\xb6\x92\x2d\xba\x6c\x94\x27\xfc\x85\x46\xd2\xef\xa2\x76\x42\x87\xb9\x6a\x25\xd5\x5f\xdb\xd2\xdc\xb8\xa8\x55\xc9\x85\x6c\x8d\xed\xa4\x0e\x64\xe6\x70\x11\xd8\x85\xb0\xb6\x55\xa5\x08\xca\xe8\xc2\x4b\xb7\x50\xa5\xcc\xe9\x34\x90\x6f\x4c\xdf\x56\x31\xe8\x4c\x52\x23\x74\x95\x71\x70\x59\xe5\x74\x69\xa8\x33\x95\x9a\x2f\xd9\x2d\xfc\xdc\x09\xbf\x43\xbd\x97\x31\xda\xa1\xb5\x2c\xc8\x93\x64\xbc\xce\x4b\xa3\xe7\xaa\xee\x9d\x7c\x92\xee\xa6\x4f\xb9\xa2\xeb\x41\x74\x9d\x10\x0d\xa7\x7c\xd1\xe5\x33\xf3\x85\xf6\x29\x6d\x84\x6f\x54\x69\x9c\x2d\xac\x93\xa5\xf2\xf2\xc5\xcf\x69\x02\xc5\x2d\x3a\x31\x1e\x05\xe8\x76\x49\x5a\x86\xcf\xc6\x5d\xdd\x33\x1f\x65\x94\x5a\xa7\x16\xc0\x61\x3a\x0a\xd2\x1d\x52\x76\x8f\xd2\xd5\x8a\x81\x98\x2a\x3b\x15\x55\xe5\xa4\xf7\xb4\x5e\x8f\x8e\x27\x32\xf4\x96\x04\xf9\xa5\x2e\x81\xdf\xdc\xb4\x95\x74\x34\x77\xa6\x23\xd3\x3b\x62\x2f\x4a\xd7\x54\x29\x24\x14\x8c\x43\xf9\x86\x8a\xc5\x50\xdd\xbd\x1c\x06\x07\xd3\xd1\x01\x87\xb4\x22\x34\xf9\xc6\x01\x02\xee\x50\xba\xb1\x4c\x77\x60\x4b\x64\x3e\xa3\x6f\xc8\xef\x46\x4a\xb5\x33\xbd\xbd\x23\x19\x92\x3c\xd6\x62\x86\x36\x4f\x26\x27\x24\x6a\x6e\x25\xda\xfb\x59\xb8\x8a\x1d\x7b\x83\xf6\x87\xc0\xc7\xb1\x7a\xd4\x6a\xa5\xae\xa4\x2e\x95\xf4\xb1\x02\x7f\x9b\xa9\xf7\x4d\x3e\x5a\x4f\x07\x5f\xfb\x14\x5c\x2f\x87\x40\x6f\x4d\xaf\xab\x38\x17\xb4\xe9\xdc\xf0\xf5\x44\xcd\x49\xe8\xe5\x53\x68\xad\x1e\xc6\xe9\x02\x22\xa4\x34\x8e\x1b\x8b\x29\x24\x3e\x07\xce\xf4\x70\x0d\x35\xbe\x47\x4b\x0b\x83\x71\x2c\x6e\xb5\x32\x00\x03\xf3\xd6\x18\x9b\x1f\x41\x1a\x00\x16\x37\xe3\xc7\x50\xb2\xb3\x88\x20\x0e\xf7\x54\xad\x33\x0b\xe5\x39\xc3\xd4\x37\xb2\x6d\xb9\xe3\xba\x55\x5a\x02\xc3\xb2\xa2\xad\x15\x0c\xd6\xf4\xe8\x11\xcd\x30\x5a\xe3\x67\xd1\x09\xa5\x73\xdf\xa4\x43\x31\x80\x8a\xeb\x41\xd2\x11\x82\x77\x46\x54\x24\xda\x36\xb6\x7f\xee\x44\xcd\xbb\xe3\xa9\x91\x4e\xc6\xba\x81\xc2\x3d\x80\xf3\x5b\x48\x36\xda\x8c\x0b\xcf\xdb\xad\x75\x44\x84\x2b\x1f\x25\xd7\x4e\x22\xca\x7a\xfd\xdd\x0c\x4e\xb5\x0f\x9c\xc0\xac\x57\x58\x46\xa9\x17\xca\x19\xcd\x56\xff\xb7\xf2\x6d\x5f\x3a\x65\xc3\x14\x6b\x9e\xc0\x77\x92\xdc\x11\xa0\x27\xaf\x5e\x4d\x8e\x2e\x4e\xdf\x5f\x26\x5b\x0f\x8a\x99\xd2\x05\x43\x93\x24\x5e\x06\xca\x98\x70\x7a\x3d\x1e\xa5\x73\xf2\x8b\x8a\x47\xab\xac\x9c\x0b\xd5\x8e\xe2\xe0\x44\x89\x89\xc1\xc9\xb8\x27\x4f\x69\xc5\xc3\xdc\x9a\x52\xb4\x98\xc6\xde\x95\x92\x49\x60\x7f\xfb\xd9\xad\x98\xf3\xd2\x66\x7f\x7b\x97\x45\xb2\x6c\x0c\xa5\xc7\x17\x17\xe7\x17\x24\x02\x6d\xaf\x6e\x8d\xd6\x7b\xdb\xab\x41\x77\xfd\x92\xde\x09\xec\x7c\x6b\x6a\xbf\xc7\x9d\xc2\x6a\x48\x4b\x3c\x4d\xbc\x7f\xae\xc0\x45\xe1\x97\x1e\x3f\x74\x4d\x21\xe6\xa6\x69\xf7\xa7\x64\x9d\x20\x3b\x4b\x8f\x63\x72\x94\x6e\xaf\x5e\x1f\x4e\x4e\xa6\x93\xf3\x0f\x17\x47\xc7\xeb\x94\x05\xef\x4e\xcf\x8e\xcf\xce\xd7\xe9\x63\x42\x0e\x49\x62\x24\x97\x80\x8b\x5f\x52\xda\x3d\x78\xf4\x0c\xee\xe0\xb4\xc6\xe4\x65\x61\x88\x77\x40\x05\xda\x59\xe8\xbe\x6d\x5f\xd2\x3a\x31\x6d\x34\x18\xca\xf8\x83\x35\xfe\x22\x18\xf3\x15\x88\xf1\x37\x71\x25\x09\xa8\x61\x2f\x43\x83\xf2\x3e\x8e\xab\x4c\xd8\xbc\x8f\x54\x1b\xac\xe4\x40\x26\x6d\xe4\x12\xa6\x4d\x30\x1e\x0b\xe2\x70\x0f\x5e\x31\xba\x37\x54\x41\x07\x88\xdf\x98\x4e\x6e\x24\x45\xce\x1d\x73\x25\x47\x3b\x1a\xb7\x94\xd7\x9f\xe9\x21\x8e\x21\x70\x43\xf6\x28\x42\xe9\x04\x7b\xfb\x60\x80\x2e\xfd\xe0\xe5\x9b\xb3\x09\x3a\x9c\x52\x21\x43\x59\x20\x21\xfe\xaf\xa6\xc3\x50\xd1\xc1\x9d\x2a\x91\x96\x4e\x36\xad\xba\x63\x78\x4d\xbe\x07\x89\x07\x29\x29\x13\xff\xe5\x06\x0e\x8c\x1c\x0c\xc6\x57\x86\x41\x20\x10\x70\x10\x2e\x24\x73\x05\xf0\x5b\x4a\x0f\xab\x48\x64\xc2\x06\x5c\x59\xe3\x15\x78\x96\x99\x0b\xcf\x0f\xf5\x96\x69\x43\xd7\x79\x9e\xa7\x89\xfc\x62\x8d\x0b\xf4\xe6\xf8\xf5\xe9\xe1\xd9\xf4\xed\xc5\xf9\xd9\xe5\xf1\xd9\x9b\x7d\x6d\xb4\x62\x1e\x11\x65\x50\x0b\x99\x6c\x42\xc2\x5f\x06\x66\x1c\x5c\x20\xdd\xe5\x37\x37\x6a\x5c\xb4\x6c\x49\x76\x19\x1a\x90\x93\x37\xf3\x00\x6a\x94\x19\xb6\xcb\x4a\x17\x38\x8d\xef\xc8\xb2\xd2\x74\x1d\x36\x8f\x1d\x61\xda\xb4\xe7\xbc\xb2\x26\x04\xeb\x6f\x83\x54\x55\xc6\xf7\x37\x25\x2d\x63\x1c\x2b\xf6\x8c\xc6\x1b\xf4\x77\x61\x1b\xfb\x3c\x7b\x9e\xbf\x48\xf8\x11\x92\x9d\xe7\x69\xc0\xab\xeb\xe4\xa7\x1e\xf3\x51\xf1\x27\x5e\xc1\x05\xf3\x49\xda\x6b\xd1\x73\x4b\x02\xbf\xda\xb8\xb3\xa2\xbc\x02\x7b\xfb\x14\xc6\x71\xd0\x3d\xe8\x8a\x5f\xf8\xaf\x6b\xcb\xef\x15\x7d\x25\x97\x23\x1c\x3f\x40\x29\xf6\x64\xe4\x20\xee\xcb\xfb\x93\xf7\xb1\x15\xbe\xb7\x5c\x66\x7c\x68\xc6\xf0\xb1\x2b\x3f\x42\x15\x25\xd2\x9f\x18\x83\xd9\x3f\x8e\x6a\x2c\x45\x27\x5d\x89\x61\x15\x23\xb9\x65\x78\x8a\xb9\x28\x7c\xb3\x16\xae\x86\x43\x84\xa6\x2b\xdd\x12\x23\x31\x9c\x97\xfe\x53\x3b\x1c\xe7\xb6\x1b\x0e\x75\x35\xfc\x32\x95\x32\x67\x0c\x5f\xb6\x86\xe6\x37\x35\x1c\x99\x0e\x5d\x90\x2e\x26\xcc\xcb\x15\x3a\x9b\xc4\x70\x99\x9f\x50\x6c\xdc\x5e\x51\x20\xfb\x72\xa3\x68\x5c\x5d\x8c\x95\x60\x97\xae\xd9\xf9\x4d\xa5\xdd\x82\x6e\xf4\x6c\x23\x1c\x15\xbd\x67\x42\x02\xd3\x45\x42\xdd\x5c\x26\x23\xcf\xfe\x1b\x00\x00\xff\xff\xb9\x3a\x40\xb6\xdd\x09\x00\x00"

func dataCommonDevVagrantfileTplBytes() ([]byte, error) {
	return bindataRead(
		_dataCommonDevVagrantfileTpl,
		"data/common/dev/Vagrantfile.tpl",
	)
}

func dataCommonDevVagrantfileTpl() (*asset, error) {
	bytes, err := dataCommonDevVagrantfileTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/common/dev/Vagrantfile.tpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"data/aws-simple/build/build-php.sh": dataAwsSimpleBuildBuildPhpSh,
	"data/aws-simple/build/template.json.tpl": dataAwsSimpleBuildTemplateJsonTpl,
	"data/aws-simple/deploy/main.tf.tpl": dataAwsSimpleDeployMainTfTpl,
	"data/common/dev/Vagrantfile.tpl": dataCommonDevVagrantfileTpl,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"data": &bintree{nil, map[string]*bintree{
		"aws-simple": &bintree{nil, map[string]*bintree{
			"build": &bintree{nil, map[string]*bintree{
				"build-php.sh": &bintree{dataAwsSimpleBuildBuildPhpSh, map[string]*bintree{
				}},
				"template.json.tpl": &bintree{dataAwsSimpleBuildTemplateJsonTpl, map[string]*bintree{
				}},
			}},
			"deploy": &bintree{nil, map[string]*bintree{
				"main.tf.tpl": &bintree{dataAwsSimpleDeployMainTfTpl, map[string]*bintree{
				}},
			}},
		}},
		"common": &bintree{nil, map[string]*bintree{
			"dev": &bintree{nil, map[string]*bintree{
				"Vagrantfile.tpl": &bintree{dataCommonDevVagrantfileTpl, map[string]*bintree{
				}},
			}},
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

