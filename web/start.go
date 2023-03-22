package web

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("/static", true)))

	tmpl, err := template.ParseFiles("web/templates/home.html")
	if err != nil {
		log.Fatal(err)
	}

	logFile, err := os.Create(".log")
	if err != nil {
		log.Fatal(err)
	}
	router.Use(gin.LoggerWithWriter(logFile))

	handlerFunction := func(c *gin.Context) {
		c.Render(http.StatusOK, render.HTML{
			Template: tmpl,
			Data:     nil,
		})
	}

	router.GET("/", handlerFunction)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
