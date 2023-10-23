package router

import (
	"github.com/gin-gonic/gin"
	"github.com/renatocardosoalves/api-go-gin/controller"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/students", controller.GetAll)
	r.GET("/students/:id", controller.GetByID)
	r.POST("/students", controller.Create)
	r.DELETE("/students/:id", controller.Delete)
	r.PUT("/students/:id", controller.Update)

	r.Run()
}
