package web

import (
	"html/template"
	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func Start() {
	router := gin.Default()

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	handlerFunction := func(c *gin.Context) {
		c.Render(http.StatusOK, render.HTML{
			Template: tmpl,
			Data:     nil,
		})
	}

	router.GET("/", handlerFunction)

	router.Run(":8080")
}
