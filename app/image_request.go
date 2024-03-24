package app

import (
	"img-to-ascii/imgmodifier"
	"io"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func processImageRequest(context *gin.Context) {
	resizeScale := parseResizeScale(context)
	imgBuf := bodyPreprocessor(bodyReader(context))

	imgmodifier.ImgToAscii(imgBuf, resizeScale)
}

func parseResizeScale(context *gin.Context) float64 {
	resizeScale, err := strconv.ParseFloat(context.Request.URL.Query().Get("resize_scale"), 64)
	// TODO: Possible errors?
	if err != nil {
		resizeScale = 1.0
	}

	return resizeScale
}

func bodyPreprocessor(requestBody io.ReadCloser) []byte {
	imgBuf, err := ioutil.ReadAll(requestBody)
	defer requestBody.Close()

	if err != nil {
		panic(err)
	}

	return imgBuf
}

func bodyReader(context *gin.Context) io.ReadCloser {
	if isMultipart(context.ContentType()) {
		mpf, err := context.MultipartForm()
		if err != nil {
			panic(err)
		}

		// If just click "Submit" it will fail here
		// Because the array is empty.
        // TODO: Fix.
		f, err := mpf.File["img"][0].Open()
		defer f.Close()
		if err != nil {
			panic(err)
		}

		return f
	}

	return context.Request.Body
}

// Requests from front-end form.
func isMultipart(requestContentType string) bool {
	return strings.Contains(requestContentType, "multipart/form-data")
}
