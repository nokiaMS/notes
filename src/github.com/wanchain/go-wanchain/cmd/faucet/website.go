// Code generated by go-bindata.
// sources:
// faucet.html
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

var _faucetHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x5a\x71\x73\xdb\xb6\x92\xff\xdb\xf9\x14\x5b\x5e\xfc\x24\x9d\x4d\x52\xb6\x93\x3c\x9f\x44\xaa\x93\x97\xf7\x5e\x2f\x37\x77\x6d\xa7\x4d\xa7\xf7\xa6\xaf\x73\x03\x12\x2b\x11\x31\x08\xb0\x00\x28\x59\xf5\xe8\xbb\xdf\x00\x20\x29\x52\x92\xdd\xa4\xc9\xdd\x34\x7f\x38\x24\xb0\xd8\x5d\xec\xfe\x16\xbb\x58\x2a\xf9\xe2\xaf\xdf\xbc\x79\xf7\x8f\x6f\xff\x06\x85\x29\xf9\xe2\x59\x62\xff\x03\x4e\xc4\x2a\x0d\x50\x04\x8b\x67\x67\x49\x81\x84\x2e\x9e\x9d\x9d\x25\x25\x1a\x02\x79\x41\x94\x46\x93\x06\xb5\x59\x86\xb7\xc1\x7e\xa2\x30\xa6\x0a\xf1\x97\x9a\xad\xd3\xe0\xbf\xc3\x1f\x5e\x87\x6f\x64\x59\x11\xc3\x32\x8e\x01\xe4\x52\x18\x14\x26\x0d\xde\xfe\x2d\x45\xba\xc2\xde\x3a\x41\x4a\x4c\x83\x35\xc3\x4d\x25\x95\xe9\x91\x6e\x18\x35\x45\x4a\x71\xcd\x72\x0c\xdd\xcb\x25\x30\xc1\x0c\x23\x3c\xd4\x39\xe1\x98\x5e\x05\x8b\x67\x96\x8f\x61\x86\xe3\xe2\xe1\x21\xfa\x1a\xcd\x46\xaa\xbb\xdd\x6e\x06\xaf\x6b\x53\xa0\x30\x2c\x27\x06\x29\xfc\x9d\xd4\x39\x9a\x24\xf6\x94\x6e\x11\x67\xe2\x0e\x0a\x85\xcb\x34\xb0\xaa\xeb\x59\x1c\xe7\x54\xbc\xd7\x51\xce\x65\x4d\x97\x9c\x28\x8c\x72\x59\xc6\xe4\x3d\xb9\x8f\x39\xcb\x74\x6c\x36\xcc\x18\x54\x61\x26\xa5\xd1\x46\x91\x2a\xbe\x89\x6e\xa2\x3f\xc7\xb9\xd6\x71\x37\x16\x95\x4c\x44\xb9\xd6\x01\x28\xe4\x69\xa0\xcd\x96\xa3\x2e\x10\x4d\x00\xf1\xe2\xf7\xc9\x5d\x4a\x61\x42\xb2\x41\x2d\x4b\x8c\x5f\x44\x7f\x8e\xa6\x4e\x64\x7f\xf8\x69\xa9\x56\xac\xce\x15\xab\x0c\x68\x95\x7f\xb0\xdc\xf7\xbf\xd4\xa8\xb6\xf1\x4d\x74\x15\x5d\x35\x2f\x4e\xce\x7b\x1d\x2c\x92\xd8\x33\x5c\x7c\x12\xef\x50\x48\xb3\x8d\xaf\xa3\x17\xd1\x55\x5c\x91\xfc\x8e\xac\x90\xb6\x92\xec\x54\xd4\x0e\x7e\x36\xb9\x8f\xf9\xf0\xfd\xa1\x0b\x3f\x87\xb0\x52\x96\x28\x4c\xf4\x5e\xc7\xd7\xd1\xd5\x6d\x34\x6d\x07\x8e\xf9\x3b\x01\xd6\x69\x56\xd4\x59\xb4\x46\x65\x91\xcb\xc3\x1c\x85\x41\x05\x0f\x76\xf4\xac\x64\x22\x2c\x90\xad\x0a\x33\x83\xab\xe9\xf4\x7c\x7e\x6a\x74\x5d\xf8\x61\xca\x74\xc5\xc9\x76\x06\x4b\x8e\xf7\x7e\x88\x70\xb6\x12\x21\x33\x58\xea\x19\x78\xce\x6e\x62\xe7\x64\x56\x4a\xae\x14\x6a\xdd\x08\xab\xa4\x66\x86\x49\x31\xb3\x88\x22\x86\xad\xf1\x14\xad\xae\x88\x38\x5a\x40\x32\x2d\x79\x6d\xf0\x40\x91\x8c\xcb\xfc\xce\x8f\xb9\x68\xee\x6f\x22\x97\x5c\xaa\x19\x6c\x0a\xd6\x2c\x03\x27\x08\x2a\x85\x0d\x7b\xa8\x08\xa5\x4c\xac\x66\xf0\xaa\x6a\xf6\x03\x25\x51\x2b\x26\x66\x30\xdd\x2f\x49\xe2\xd6\x8c\x49\xec\x0f\xae\x67\x67\x49\x26\xe9\xd6\xf9\x90\xb2\x35\xe4\x9c\x68\x9d\x06\x07\x26\x76\x07\xd2\x80\xc0\x9e\x43\x84\x89\x76\x6a\x30\xa7\xe4\x26\x00\x27\x28\x0d\xbc\x12\x61\x26\x8d\x91\xe5\x0c\xae\xac\x7a\xcd\x92\x03\x7e\x3c\xe4\xab\xf0\xea\xba\x9d\x3c\x4b\x8a\xab\x96\x89\xc1\x7b\x13\x3a\xff\x74\x9e\x09\x16\x09\x6b\xd7\x2e\x09\x2c\x49\x98\x11\x53\x04\x40\x14\x23\x61\xc1\x28\x45\x91\x06\x46\xd5\x68\x71\xc4\x16\xd0\x3f\xfe\x1e\x39\xfd\x8a\xab\x56\xaf\x98\xb2\x75\xb3\xad\xde\xe3\xc1\x0e\x1f\xdf\xc4\x2d\x34\x0f\x72\xb9\xd4\x68\xc2\xde\x9e\x7a\xc4\x4c\x54\xb5\x09\x57\x4a\xd6\x55\x37\x7f\x96\xb8\x51\x60\x34\x0d\x6a\xc5\x83\xe6\xf8\x77\x8f\x66\x5b\x35\xa6\x08\xba\x8d\x4b\x55\x86\xd6\x13\x4a\xf2\x00\x2a\x4e\x72\x2c\x24\xa7\xa8\xd2\xe0\x7b\x99\x33\xc2\x41\xf8\x3d\xc3\x0f\xdf\xfd\x27\x34\x2e\x63\x62\x05\x5b\x59\x2b\xf8\x91\x88\xbc\x20\x4c\x00\xa1\xd4\xc2\x35\x8a\xa2\x9e\x22\x0e\xbb\xc7\xaa\x86\x99\x11\x7b\xaa\xb3\x24\xab\x8d\x91\x1d\x61\x66\x04\x64\x46\x84\x14\x97\xa4\xe6\x06\xa8\x92\x15\x95\x1b\x11\x1a\xb9\x5a\xd9\x4c\xe7\x37\xe1\x17\x05\x40\x89\x21\xcd\x54\x1a\xb4\xb4\xad\x0f\x89\xae\x64\x55\x57\x8d\x17\xfd\x20\xde\x57\x44\x50\xa4\xd6\xe7\x5c\x63\xb0\xf8\x8a\xad\x11\x4a\xb4\x7b\x39\x3b\x04\x44\x4e\x14\x9a\xb0\xcf\xf2\x08\x16\x49\xec\x55\xf1\x1b\x82\xe6\x5f\x52\xf3\x96\x53\xb7\x81\x12\x45\x0d\x83\xb7\x50\xd9\x53\x25\x58\x3c\x3c\x28\x22\x56\x08\xcf\x19\xbd\xbf\x84\xe7\xa4\x94\xb5\x30\x30\x4b\x21\x7a\xed\x1e\xf5\x6e\x37\xe0\x0e\x90\x70\xb6\x48\xc8\x53\xe0\x06\x29\x72\xce\xf2\xbb\x34\x30\x0c\x55\xfa\xf0\x60\x99\xef\x76\x73\x78\x78\x60\x4b\x78\x1e\x7d\x87\x39\xa9\x4c\x5e\x90\xdd\x6e\xa5\xda\xe7\x08\xef\x31\xaf\x0d\x8e\x27\x0f\x0f\xc8\x35\xee\x76\xba\xce\x4a\x66\xc6\xed\x72\x3b\x2e\xe8\x6e\x67\x75\x6e\xf4\xdc\xed\x20\xb6\x4c\x05\xc5\x7b\x78\x1e\x7d\x8b\x8a\x49\xaa\xc1\xd3\x27\x31\x59\x24\x31\x67\x8b\x66\xdd\xd0\x48\x71\xcd\xf7\x68\x89\x2d\x5c\x3a\x94\xbb\xa0\x71\xaa\xf6\x35\x3d\x11\x03\xab\xb0\xd3\xbe\x41\x83\x66\x06\xef\x70\x9b\x06\x0f\x0f\xfd\xb5\xcd\x6c\x4e\x38\xcf\x88\xb5\x8b\xdf\x5a\xb7\xe8\x57\xb4\x28\x5d\x33\xed\x0a\xaa\x45\xab\xc1\x5e\xed\x0f\x0c\xea\x83\x63\xcb\xc8\x6a\x06\x37\xd7\xbd\x33\xeb\x54\xbc\xbf\x3a\x88\xf7\x9b\x93\xc4\x15\x11\xc8\xc1\xfd\x0d\x75\x49\x78\xfb\xdc\xc4\x4a\x2f\xf4\x0e\x17\x85\xf6\x84\xee\x54\xeb\x4e\xfa\xe9\x1c\xe4\x1a\xd5\x92\xcb\xcd\x0c\x48\x6d\xe4\x1c\x4a\x72\xdf\x65\xbb\x9b\xe9\xb4\xaf\xb7\x2d\x04\x49\xc6\xd1\x9d\x2d\x0a\x7f\xa9\x51\x1b\xdd\x9d\x24\x7e\xca\xfd\xb5\x07\x0a\x45\xa1\x91\x1e\x58\xc3\x4a\xb4\xa6\x75\x54\x3d\xd7\x77\xc6\x3c\xa9\xfb\x52\xca\x2e\x81\xf4\xd5\x68\x58\xf7\x72\x5d\xb0\x48\x8c\xda\xd3\x9d\x25\x86\x7e\x54\x02\x50\xb6\xc0\x7b\xec\xfc\xf7\xe7\x99\xdd\x7b\x85\xa8\x7c\x75\x61\x21\x0b\xee\x35\x89\x0d\xfd\x04\xc9\x16\x84\x19\xd1\xf8\x21\xe2\x5d\x9e\xdf\x8b\x77\xaf\x9f\x2a\xbf\x40\xa2\x4c\x86\xc4\x7c\x88\x02\xcb\x5a\xd0\xde\xfe\x7f\x24\xe2\x53\xc5\xd7\x82\xad\x51\x69\x66\xb6\x1f\x2a\x1f\xe9\x5e\x01\xff\x3e\x54\x21\x89\x8d\x7a\x1a\x69\xfd\x97\xcf\x14\xda\xbf\x55\x8e\xdc\x2c\xfe\x5d\x6e\x80\x4a\xd4\x60\x0a\xa6\xc1\x26\xd6\x2f\x93\xb8\xb8\xe9\x48\xaa\xc5\x3b\x3b\xd1\x25\xd6\xa5\xab\x2c\x80\x69\x50\xb5\x70\x89\x57\x0a\x30\x05\x0e\xab\x91\x26\x47\x47\xf0\x4e\xda\x8a\x6e\x8d\xc2\x40\x49\x38\xcb\x99\xac\x35\x90\xdc\x48\xa5\x61\xa9\x64\x09\x78\x5f\x90\x5a\x1b\xcb\xc8\x9e\x1f\x64\x4d\x18\x77\xc1\xe4\x7c\x0a\x52\x01\xc9\xf3\xba\xac\x6d\x45\x2a\x56\x80\x42\xd6\xab\xc2\xaa\x03\x46\x82\xcf\x4b\x5c\x8a\x55\xa7\x8d\xae\x48\x09\xc4\x18\x92\xdf\xe9\x4b\x68\x0f\x05\x20\x0a\xc1\x30\xa4\x76\x55\x2e\xcb\x52\x0a\xb8\x51\x14\x2a\xa2\xcc\x16\xf4\xb0\xb0\x20\x79\xee\x92\x5c\x04\xaf\xc5\x56\x0a\x84\x82\xac\x9d\x7e\xf0\xce\xdf\x25\x2e\xe1\x2b\x29\x57\x1c\x2f\xac\x7a\x7f\x27\x39\x66\x52\x76\xcb\xa0\x24\xdb\x56\x6e\xb3\x89\x0d\x33\x05\xf3\x56\xaa\x50\x95\x96\x07\x05\xce\x4a\x66\x74\x94\xc4\xd5\xfe\x64\xdd\xe7\x68\x1e\x16\x52\xb1\x5f\x6d\x79\xc3\xfb\xc7\xa8\x39\x38\x64\xda\x33\xd2\xf9\x9f\xe3\xd2\xcc\xe0\x85\x3f\x23\x0f\x11\xdd\xdc\x83\x4e\xc1\xb9\xe5\xe9\xee\x97\x36\xf1\xcc\xe0\xc6\x17\xb5\xbe\xa0\xa0\xa6\xa7\x01\x3d\x00\x9d\x17\x7a\x7b\x5b\xdd\x77\x7a\x74\x95\xf1\xb4\x63\x62\x81\x30\x34\xca\x9a\xf5\xec\x59\x92\x3b\x04\x02\x09\x39\xb8\x27\x37\x4a\xbb\x5b\x16\x73\x5d\x82\xd8\x6c\x10\xcd\x97\x36\x88\xd3\xef\x3c\x43\x26\x56\xe7\xd7\x53\x0f\x4c\xfb\x60\xd9\x9f\x5f\x4f\x99\x30\xf2\xfc\x7a\x3a\xbd\x9f\x7e\xe0\xbf\xf3\xeb\xa9\x14\xe7\xd7\x53\x53\xe0\xf9\xf5\xf4\xfc\xfa\xa6\x0f\x69\x3f\xd2\x86\x81\xa5\x42\x6d\xa5\xb5\x48\x0f\xc0\x10\xb5\x42\x93\x06\xff\x43\x32\x59\x9b\x59\xc6\x89\xb8\x0b\x16\x4e\x5d\x5b\x75\x38\x14\x9c\xae\x52\xa1\x22\xda\x42\xc2\x6a\xec\x50\xd2\x74\x44\x34\x8c\x75\xad\x94\xac\x85\xcd\x8e\x60\xf7\xec\x62\x55\x8c\x2c\xca\xac\x61\x26\x51\x92\xa9\x78\xf1\x46\x56\xdb\xd0\x31\x71\xcb\x8f\xcc\xa8\xeb\xaa\x92\xca\x44\x7d\x73\x12\x7b\x1b\xe2\xa8\xe3\xdb\xe9\xcb\xdb\x57\x4f\xaa\xaf\x6d\xad\xed\xf6\xd0\x69\x48\x32\xb9\x46\xf0\x95\x7d\x26\xef\x81\x08\x0a\x4b\xa6\x10\xc8\x86\x6c\xbf\x48\x62\xea\xee\x61\x9f\x8e\xda\x95\x0b\xb4\xb0\xe2\xb5\xb6\xa5\x08\xb3\x81\xfa\x87\x82\xb0\x3f\x09\xe0\x5b\x5e\xeb\x4b\xa8\xea\x8c\x33\x5d\x00\x01\x81\x1b\x48\xb4\x51\x52\xac\x16\x6e\x34\xb7\xf7\x54\xf7\x0a\x95\xd4\xe6\x29\x34\x60\x99\x21\xa5\x27\xf0\xf0\x3b\xe1\x60\xe5\x39\x17\xfe\xff\xbb\x6f\xd9\x1c\x8e\x7f\x28\x97\xb5\x27\xf6\x1f\xd5\x5f\x47\xe1\xbb\xd9\x6c\xa2\xd6\x92\x2e\x76\x0b\xe4\x55\x6c\xd3\x58\x2d\x98\xd9\xc6\xfe\x14\x94\x22\xfe\x92\xd1\xf4\xfa\xf6\xfa\xd5\xab\xeb\x17\xff\x76\xfb\xf2\xe5\xf5\xed\x8b\x97\x8f\x05\x76\x07\x8a\xdf\x1f\xd7\xfe\x36\xf4\xb5\x7c\x5d\x9b\xa2\xbb\x0a\x79\xbc\xb4\x25\xb8\x2d\xb5\xa8\xbd\x4a\xaa\xe0\x77\x63\xa8\x16\xb6\x9e\x0c\x09\x3f\x59\x0a\x7e\x04\x8a\x1c\x8c\x9e\xd0\xec\x13\xa1\xd5\xc2\xc7\x22\x45\xd6\xc6\xee\xb0\xed\xc8\x30\x29\x3a\x38\x5d\x82\x66\x65\xc5\xb7\x90\xef\xbd\x7e\x1a\x57\x8f\x3a\xe5\x37\x61\x35\x74\x9b\x07\x99\x2b\xe3\x4a\x49\xd1\xd6\x6e\xba\xd6\x39\x56\xae\x55\x6f\x2b\xa2\xbf\x6c\x7f\x25\xc2\x30\x81\x6d\xe5\x14\xc1\x37\x82\x6f\xa1\xd6\x08\x4b\xa9\x80\x62\x56\xaf\x56\xae\xd8\x53\x50\x29\xb6\x26\x06\xdb\x72\x49\x37\xa8\xe8\x40\xd1\xbb\xa0\xda\xda\x95\xf7\x4a\xc9\x7f\xc8\x1a\x72\x5b\xba\x29\x92\xdf\xf9\x48\xa9\x95\xb2\x91\x52\xa1\xdf\x4d\x57\xb0\x65\xc8\xe5\xc6\x91\xf8\x7d\x2f\x19\x72\x57\xbd\x69\x44\x28\xe4\x06\xca\x3a\x77\x01\x69\xab\x33\xb7\x89\x0d\x61\x06\x6a\x61\x18\xf7\xf6\x34\xb5\x12\xb6\xd6\xc3\x41\x91\x75\x74\x85\x4f\xb0\x5c\xbc\x2b\xf0\x44\x61\xdb\x5d\xbe\x41\xe1\x1b\x4f\x0e\x95\x92\x06\x73\xeb\x50\x20\x2b\xc2\x84\xb6\x1e\x71\x65\x1c\x96\x1f\x70\x39\xef\x9e\x9a\x87\x7d\x9b\xd9\x4d\xc7\x31\x7c\xc5\x65\x46\x38\xac\x2d\xd2\x33\x6e\xeb\x72\x09\x85\xb4\x5b\xef\x59\x4b\x1b\x62\x6a\x0d\x72\xe9\x46\xbd\xe6\x76\xfd\x9a\x28\xeb\x41\x2c\x2b\x03\x69\xd3\x24\xb5\x63\x1a\xd5\xba\x69\xfd\xda\x57\xc3\x50\x0d\xe6\x3b\xab\xa7\xf0\xd3\xcf\xf3\x67\x8d\x2a\x7f\xc5\xa5\x83\x84\xc5\xb7\xdf\xb2\x29\x88\x81\x5c\x21\x31\xa8\x21\xe7\x52\xd7\xca\x6b\x48\x95\xac\xc0\x6a\xd9\x72\x6a\x39\xdb\x89\xca\x49\x6b\x99\x8c\x0b\xa2\x8b\x49\xd3\xe3\x55\xe8\xbc\xd4\xcd\xb5\xe3\x67\x16\x75\x63\xcb\x80\xa5\xd3\x39\xb0\xa4\xe5\x1b\x71\x14\x2b\x53\xcc\x81\x5d\x5c\x74\xc4\x67\x6c\x09\xe3\x96\xe2\x27\xf6\x73\x64\xee\x23\x2b\x05\xd2\x14\xfa\xd2\x9c\xc0\x86\x8f\xae\x38\xcb\x71\xcc\x2e\xe1\x6a\x32\x6f\x67\x33\x85\xe4\xae\x7d\x6b\xfc\xe8\xff\x73\x7f\x77\xf3\xa1\x65\x9c\xf1\x07\xb6\xf1\x2d\x1c\x0d\x04\x56\x4c\x1b\xa8\x15\x87\x26\x86\xbd\x0b\x3a\x87\x38\xba\xbe\x55\x8e\x70\xd9\x3c\x34\x98\x6a\xb7\xe0\xd9\x44\x1a\x05\x1d\xff\xc7\xf7\xdf\x7c\x1d\x69\xa3\x98\x58\xb1\xe5\x76\xfc\x50\x2b\x3e\x83\xe7\xe3\xe0\x5f\x6a\xc5\x83\xc9\x4f\xd3\x9f\xa3\x35\xe1\x35\x5e\x3a\x7f\xcf\xdc\xdf\x23\x29\x97\xd0\x3c\xce\x60\x28\x70\x37\x99\xcc\x4f\xb7\xbb\x7a\xdd\x39\x85\x1a\xcd\xd8\x12\x76\xc0\x3f\xb4\x11\x81\x12\x4d\x21\x5d\xe8\x2a\xcc\xa5\x10\x98\x1b\xa8\x2b\x29\x1a\x93\x00\x97\x5a\xef\x81\xd8\x52\xa4\xc7\xa0\x68\xe8\x53\x97\xac\x7f\xc4\xec\x7b\x99\xdf\xa1\x19\x8f\xc7\x1b\x26\xa8\xdc\x44\x5c\xfa\xa3\x36\xb2\x41\x2a\x73\xc9\x21\x4d\x53\x68\xb2\x68\x30\x81\x2f\x21\xd8\x68\x9b\x4f\x03\x98\xd9\x47\xfb\x34\x81\x0b\x38\x5c\x5e\xd8\x7c\x7f\x01\x41\x4c\x2a\x16\x4c\x7c\x38\xb4\x86\x97\xa2\x44\xad\xc9\x0a\xfb\x0a\xba\xfb\x6d\x07\x32\xbb\x8f\x52\xaf\x20\x05\xe7\xa0\x8a\x28\x8d\x9e\x24\xa2\xc4\x90\x16\x6d\x16\xb3\x8e\x2c\x4d\x41\xd4\x9c\xef\x41\xea\x83\x62\xde\xc2\x6f\x40\x1e\xf9\x5c\xf3\x45\x9a\x42\x2d\xa8\x33\x31\xdd\xaf\xb4\xce\xf7\xbd\x90\x49\x64\xf3\xc2\x7e\xc5\x64\xde\x47\xf3\x80\x1b\xd2\xdf\x62\x87\xf4\x90\x1f\xd2\x47\x18\xba\xd6\xd3\x53\xfc\x7c\xab\xaa\xc7\xce\x0d\x3c\xc2\x4d\xd4\x65\x86\xea\x29\x76\xbe\xf5\xd4\xb0\x73\xa6\x7e\x2b\x4c\x6f\xed\x25\x5c\xbd\x9a\x3c\xc2\x1d\x95\x92\x8f\x32\x17\xd2\x6c\xc7\x0f\x9c\x6c\x6d\xcd\x04\x23\x23\xab\x37\xae\x57\x34\xba\x74\x19\x77\x06\x1d\x87\x4b\xf7\x05\x60\x06\x23\xf7\x66\xe7\x59\x89\x6e\xd5\xcb\xe9\x74\x7a\x09\xed\xa7\xb3\xbf\x10\x1b\x84\xaa\xc6\xdd\x23\xfa\xe8\x3a\xcf\x6d\xde\xff\x14\x8d\x1a\x1e\x9d\x4e\xcd\xfb\x27\x68\xd5\xe5\x86\x81\x5a\xf0\xa7\x3f\xc1\xd1\xec\x10\xc6\x71\x0c\xff\x45\xd4\x9d\x6b\xeb\x54\x0a\xd7\xae\xf5\xd3\xd1\x97\x4c\x6b\xd7\x54\xd1\x40\xa5\xc0\x66\xcd\xc7\x1d\xfb\x47\x3a\x36\x64\xb0\x80\xe9\xa1\x82\xf6\x38\xec\xa5\x85\x13\xd9\xa2\xc7\x77\x98\x08\xce\x76\x7d\x79\x83\x95\xac\x44\xf8\x22\x85\x20\xe8\x2f\x3e\xa2\xb0\x04\x1d\xb3\x33\x8d\xe6\x9d\xf7\xc5\xb8\xc9\x8e\xa7\x72\xd7\xe4\x12\x6e\xa6\xd3\xe9\xe4\x48\x89\xdd\xde\xbc\xaf\x2b\x5b\x36\x01\x11\x5b\x77\x24\x76\xb6\x75\x85\xa3\x2d\x81\xec\x91\xc6\x21\x97\x9c\xfb\x9a\xa5\x59\x6a\x0d\xdc\x34\xc1\x52\x08\xaf\xe6\x27\xb2\x68\xcf\x92\xbd\xad\x1d\xba\xe7\x84\xed\x0f\x5d\x34\xb4\xd9\x01\x71\x78\x35\x70\xca\xc0\x5f\xa7\x1d\x73\xd6\xe9\xcd\xf6\x16\x3d\x70\xd7\xde\x5f\x87\x36\xeb\xe9\xef\xf9\x5c\x5c\x7d\xe0\x36\xba\xe9\xaa\xd6\xc5\xf8\x40\xd1\xc9\xfc\xd8\x37\x6f\x0d\x2a\x5b\x25\x4b\x9b\xb2\xac\x2f\xec\x55\x40\xe1\x91\x4b\x5c\xa9\xae\x30\x54\x28\x28\xaa\xb6\xa4\xf0\x95\xbd\x2d\x00\x07\x2e\xf3\xb7\xca\x3e\x9c\x3e\x32\x60\x5c\x49\x26\x05\x02\x00\x1c\x04\x81\x03\xea\x00\xa9\x96\x18\x39\xa9\x34\x52\x48\xc1\xff\x92\x61\x3c\x89\x6a\xc1\xee\xc7\x93\xb0\x79\x3f\xe4\xd1\xce\xcf\xbb\x6b\x62\xab\xf6\x45\x0a\x41\x62\x14\x30\x9a\x8e\x02\xb8\x38\x15\x82\x36\xeb\x8e\x16\x7b\x0d\xfa\x4b\x01\x12\x43\x17\xae\xa1\xed\xef\x6b\xff\x0c\x32\x92\xdf\xad\xdc\x45\x68\x66\x4b\xad\xf1\x11\x5b\xb2\x26\x86\x28\xc7\x75\x32\x87\x3d\x79\x73\x51\xcc\xad\x73\xe6\xe0\x6f\xa4\xae\x6f\x0e\xdd\x97\x26\xf7\x96\x49\x45\x51\x85\x8a\x50\x56\xeb\x19\xbc\xa8\xee\xe7\xff\x6c\xbf\xc4\xb9\xee\xfe\x93\xaa\x56\x0a\x17\x47\x1a\x35\x4d\xe2\x0b\x08\x92\xd8\x12\xfc\x16\x9b\x6e\xb3\xfd\x5f\x50\xc0\x89\x6f\x18\xd0\xfd\xbe\xa1\x19\x2f\x19\xa5\x1c\xad\xc2\x7b\xf6\x36\x18\xad\xff\xfb\x21\x35\x14\x09\xcd\xc7\x8b\xfd\x9a\x1d\x20\xd7\xf8\xc4\x82\xee\x3b\xc8\xc8\x02\x20\xb4\x5b\x66\xce\xe6\xcd\x65\xdb\x0d\xab\x91\xb3\x45\xf3\x7b\x18\x5a\x2b\x57\x6b\x8d\xc3\x06\x60\x97\x30\xd2\xb6\xf6\xa3\x7a\x34\x89\x8a\xba\x24\x82\xfd\x8a\x63\x9b\x97\x26\xde\x56\xee\xc3\x4a\x70\x7c\x24\x1f\x29\xb3\xff\xe2\x31\x6a\x73\xdc\xa8\x31\xe2\xa8\xf5\xee\x8b\xfd\xdd\x7e\x06\xd3\xf9\xe8\x23\x2d\x74\x5a\x4a\x98\x11\x05\xfd\x97\xb0\x4d\xbe\xa0\xa4\x95\xde\xce\x65\x44\x8d\x7c\x27\xc3\xd5\xe7\x42\x6e\xd2\xd1\xcd\xb4\x53\xd2\x3b\xda\xf9\x79\xd4\x60\xed\xc8\x19\x56\xcb\x36\x34\x17\x70\x33\xfd\x1c\xda\xfa\x6e\xc8\xc1\x0e\x8c\x62\x15\x52\x20\xb9\x61\x6b\xfc\x3f\xd8\xc8\x67\x30\xf2\x47\xab\x68\x71\xd8\x1a\xcf\xc1\x74\xa0\xaf\x9d\xed\x6c\xfb\xaf\x36\xde\x20\x76\x16\xbe\x80\xe0\xe4\x46\x1e\x45\xe2\x01\xe1\x41\x68\x3f\x1e\xf7\xee\x4b\x61\x70\x98\x53\x6c\xb5\xdb\x7d\xe3\x9e\x44\x85\x29\xf9\x38\x48\x8c\xfb\xa5\x93\xd5\xb9\xe3\xe0\x18\xf8\xe1\x61\x49\xb7\x1b\x5e\x64\xec\xfd\x1d\x0f\xee\x59\xd0\x2b\x4e\xba\xbb\x58\x5b\x89\xc0\x6e\xff\x83\xb0\x38\x86\xef\x0d\x51\x06\x08\xfc\xf0\x16\xea\x8a\x12\x63\xb3\x97\x04\x9b\x1f\x7d\xd7\xb9\xfd\xc5\x58\x46\x94\x86\xa5\x54\x1b\xa2\x68\xd3\x9f\x31\x05\x6e\xdd\x27\xb9\xb6\xf4\xd3\x68\xde\xda\x53\x6c\x4d\xf8\xf8\xe8\xde\xf7\x7c\x3c\x8a\xfa\x2e\x1f\x4d\x22\x24\x79\x71\x4c\xe8\x32\x56\x27\x37\x85\xaf\xdd\x15\x60\xfc\x7c\x6c\x0a\xa6\x27\x11\x31\x46\x8d\x47\x03\x30\x8c\x26\xd6\xaf\x57\xbd\x2b\x59\xb7\x3c\x19\x84\xd5\x53\x3c\xf6\xc5\x74\x57\x08\xb4\xe4\xb9\xd6\x63\x8f\xab\xd1\x65\x8f\xf7\x10\x56\xa3\xf3\x51\xe7\xa8\x7d\x78\xef\xf7\x91\x9e\xd4\x64\xc0\x7a\x64\xa3\x6c\x74\x24\x9e\x50\xfa\xc6\xc6\xcf\x38\x38\x11\xe9\x87\xe8\x98\x74\xc6\xf6\xe7\xf5\x93\x56\xf6\xbf\xae\x79\xc4\xc4\x8c\x8e\x26\x91\xae\x33\xdf\x9b\x18\xbf\xec\x2e\x60\x2d\x99\x03\xef\x61\x2a\x38\x2a\x28\xac\x88\x61\x51\x11\x1e\x14\x21\x4f\x64\x8d\x46\xa4\xdf\xd5\xee\xd2\x1a\x7c\x3a\xe9\x5a\x5b\x7f\xd3\xb6\xb8\xf2\xad\xff\x0d\x66\xda\x75\x12\xa0\xc1\xbb\xeb\xe6\xf8\xae\xcd\xeb\x6f\xdf\xf6\x3a\x37\x5d\x44\x8c\x1d\xf7\xee\xc7\x9c\xa7\xfa\x24\x27\x7f\x3d\xba\xd9\x6c\x22\xff\x45\xcb\xb5\xf1\xbb\x46\x4a\x4c\x2a\x16\xbd\xd7\x01\x10\xbd\x15\x39\x50\x5c\xa2\x5a\xf4\xd8\x37\xdd\x95\x24\xf6\xbf\x6b\x4c\x62\xff\xd3\xed\xff\x0d\x00\x00\xff\xff\xbd\xed\xb8\xa0\xcb\x2d\x00\x00")

func faucetHtmlBytes() ([]byte, error) {
	return bindataRead(
		_faucetHtml,
		"faucet.html",
	)
}

func faucetHtml() (*asset, error) {
	bytes, err := faucetHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "faucet.html", size: 11723, mode: os.FileMode(511), modTime: time.Unix(1514104395, 0)}
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
	"faucet.html": faucetHtml,
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
	"faucet.html": &bintree{faucetHtml, map[string]*bintree{}},
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

