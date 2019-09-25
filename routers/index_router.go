package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/thornshell/goblogV2/controllers"
)

func Init(engine *gin.Engine) {
	engine.GET("/index", controllers.Index)
}
