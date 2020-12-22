package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/service"
	"github.com/qbbang/goGin-framework/01-getting-started/entity"
)

type VideoController interface {
	FindAll() []entitiy.Video
	Save(ctx *gin.Context) entity.Video
}

type controller struct {
	service service.VideoService
}

// 생성자
func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func (c *controller) FindAll() []entitiy.Video {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) {
	var video entity.Video
	ctx.BindJSON(&video)
	c.service.Save(video)
	return video
}
