package routes

import (
	"github.com/Yummiii/Nayu/internal/handlers"
	"github.com/gin-gonic/gin"
)

func siteRoutes(r *gin.Engine) {
	r.GET("/", handlers.IndexHandler())
	r.GET("/login", handlers.LoginHandler())
}
