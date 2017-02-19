package assets

import (
	"net/http"

	"github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
)

func FileServerHandler() echo.HandlerFunc {
	fs := &assetfs.AssetFS{
		Asset:     Asset,
		AssetDir:  AssetDir,
		AssetInfo: AssetInfo,
	}
	return echo.WrapHandler(http.FileServer(fs))
}
