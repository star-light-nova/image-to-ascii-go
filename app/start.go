package app

import (
	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()

	router.LoadHTMLFiles("./index.tmpl", "./pages/404.tmpl")

	router.POST("/", asciiImage)
	router.GET("/", mainPage)

	router.NoRoute(notFound)

	assets := router.Group("/assets/")
	assets.StaticFile("favicon.ico", "./favicon.ico")
	assets.StaticFile("styles.css", "./assets/style.css")
	assets.StaticFile("index.js", "./assets/index.js")

	router.Run()
}
