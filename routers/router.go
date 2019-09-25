package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/thornshell/goblogv2/controllers"
)

func Init(engine *gin.Engine) {
	engine.GET("/aaa", controllers.Index)
}
