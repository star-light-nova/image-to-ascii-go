package main

import (
	// "fmt"
	// "encoding/json"
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
	router.POST("/", getImage)

	router.Run()
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
		bytes := imgmodifier.ImgToAscii(imgBuf, resizeScale)

		context.JSON(http.StatusOK, toOutput(bytes))
	} else {
		imgBuf := bodyPreprocessor(body)
		bytes := imgmodifier.ImgToAscii(imgBuf, resizeScale)

		context.JSON(http.StatusOK, toOutput(bytes))
	}
}
