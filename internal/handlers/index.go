package handlers

import (
	"context"

	"github.com/Yummiii/Nayu/web/pages"
	"github.com/gin-gonic/gin"
)

func IndexHandler() gin.HandlerFunc {
	template := pages.Index()

	return func(c *gin.Context) {
		_ = template.Render(context.Background(), c.Writer)
	}
}
