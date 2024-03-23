package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type mainPageData struct {
	URL string
}

func mainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", &mainPageData{URL: "http://localhost:8080"})
}

func notFound(context *gin.Context) {
	context.HTML(http.StatusOK, "404.tmpl", nil)
}

func asciiImage(context *gin.Context) {
	processImageRequest(context)

	context.FileAttachment("./asciioutput.txt", "asciioutput.txt")
}
