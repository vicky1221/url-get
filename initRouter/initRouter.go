package initRouter

import (
	"net/http"
	"strings"
	"url-get/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/statics", "./statics")
	// router.Any("/", retHelloGinAndMethod)
	// router.GET("/", retHelloGinAndMethod)
	// router.POST("/", retHelloGinAndMethod)

	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	return router
}

func retHelloGinAndMethod(ctx *gin.Context) {
	ctx.String(http.StatusOK, "hello gin "+strings.ToLower(ctx.Request.Method)+" method")
}
