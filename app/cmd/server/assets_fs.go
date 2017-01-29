package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

type AssetFileSystem struct {
	http.FileSystem
}

type AssetFile struct {
	io.ReadSeeker
	io.Closer
	os.FileInfo
}

func FileServerHandler() echo.HandlerFunc {
	return echo.WrapHandler(http.FileServer(&AssetFileSystem{}))
}

// @param name URL パスが渡ってくる。ルーティングを go-bindata で固めたパスに合わせる
func (fs *AssetFileSystem) Open(name string) (http.File, error) {
	path := strings.TrimLeft(name, "/")
	data, err := Asset(path)
	if err != nil {
		return nil, err
	}
	info, _ := AssetInfo(path)
	file := &AssetFile{
		bytes.NewReader(data),
		ioutil.NopCloser(nil),
		info,
	}
	return file, nil
}

func (f *AssetFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, errors.New("not a directory")
}

func (f *AssetFile) Stat() (os.FileInfo, error) {
	return f, nil
}
