package routes

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Assets struct {
	Assets   fs.FS
	Dist embed.FS
}

func SetupRoutes(r *gin.Engine, assets Assets) {
	assetsFs := http.FS(assets.Assets)

	r.StaticFS("/assets", assetsFs)
	r.StaticFileFS("/favicon.ico", "favicon.ico", assetsFs)
	r.StaticFileFS("/tailwind.css", "dist/tailwind.css", http.FS(assets.Dist))

	siteRoutes(r)
}
