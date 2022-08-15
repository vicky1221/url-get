package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "download.html", gin.H{
		"title":   "上传保险精灵下载链接",
		"usednum": GetusedLinkNumber(),
		"leftnum": GetusedLinkNumber(),
	})
}
