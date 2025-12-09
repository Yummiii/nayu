package main

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/Yummiii/Nayu/internal/routes"
	"github.com/gin-gonic/gin"
)

//go:embed web/assets/*
var assetsEmbed embed.FS

//go:embed dist/tailwind.css
var dist embed.FS

func main() {
	assets, err := fs.Sub(assetsEmbed, "web/assets")

	if err != nil {
		fmt.Println("Error loading assets:", err)
		return
	}

	router := gin.Default()

	routes.SetupRoutes(router, routes.Assets{
		Assets:   assets,
		Dist: dist,
	})

	err = router.Run(":8080")

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
