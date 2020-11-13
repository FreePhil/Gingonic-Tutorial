package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"phil.com/gingonic/controller"
	"phil.com/gingonic/middleware"
	"phil.com/gingonic/service"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService service.VideoService = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {

	// setup middle ware for recovery and logging
	//
	setupLogOutput()
	server := gin.New()

	server.Static("/css", "./template/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middleware.Logger(),
		middleware.BasicAuth(), gindump.Dump())

	apiRoutes := server.Group("/api")

	// routing
	//
	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}

func setupLogOutput() {

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}