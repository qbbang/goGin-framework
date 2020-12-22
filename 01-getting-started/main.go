package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/service"
	"github.com/qbbang/goGin-framework/01-getting-started/controller"
)

var (
	videoService service.VideoService       = service.New()
	videoService controller.VideoController = controller.New(videoService)
)

func main() {
	// 서버 초기화
	server := gin.Default()

	// 첫번째 엔드 포인트 - Context으로 미드뤠어간에 변수를 전달하고, 흐름을 관리, JSON
	/*
		server.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK!!",
			})
		})
	*/

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))

	})

	server.Run(":8080")

}
