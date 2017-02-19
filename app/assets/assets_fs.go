package assets

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
)

func FileSystem() http.FileSystem {
	return &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
	}
}
