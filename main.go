package main

import (
	"img-to-ascii/imgmodifier"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("./index.tmpl")

	router.POST("/", getImage)
	router.GET("/", getMainPage)

	assets := router.Group("/assets/")
	assets.StaticFile("favicon.ico", "./favicon.ico")
	assets.StaticFile("styles.css", "./assets/style.css")
	assets.StaticFile("index.js", "./assets/index.js")

	router.Run()
}

type MainPageData struct {
	URL string
}

func getMainPage(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", &MainPageData{URL: "http://localhost:8080"})
}

// Requests from front-end form.
func isMultipart(requestContentType string) bool {
	return strings.Contains(requestContentType, "multipart/form-data")
}

func toOutput(bytes [][]byte) []string {
	output := []string{}
	for _, arr := range bytes {
		output = append(output, string(arr))
	}

	return output
}

func bodyPreprocessor(requestBody io.Reader) []byte {
	imgBuf, err := ioutil.ReadAll(requestBody)
	if err != nil {
		panic(err)
	}

	return imgBuf
}

func getImage(context *gin.Context) {
	body := context.Request.Body
	defer body.Close()

	resizeScale, err := strconv.ParseFloat(context.Request.URL.Query().Get("resize_scale"), 64)
	if err != nil {
		resizeScale = 1.0
	}

	if isMultipart(context.ContentType()) {
		mpf, err := context.MultipartForm()
		if err != nil {
			panic(err)
		}

		f, err := mpf.File["img"][0].Open()
		if err != nil {
			panic(err)
		}

		imgBuf := bodyPreprocessor(f)
		imgmodifier.ImgToAscii(imgBuf, resizeScale)

		context.FileAttachment("./asciioutput.txt", "asciioutput.txt")
	} else {
		imgBuf := bodyPreprocessor(body)
		imgmodifier.ImgToAscii(imgBuf, resizeScale)

		context.FileAttachment("./asciioutput.txt", "asciioutput.txt")
	}
}
