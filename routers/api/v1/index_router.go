package v1

import (
	"goblogV2/controllers"

	"github.com/gin-gonic/gin"
)

func Init(engine *gin.Engine) {
	engine.GET("/index", controllers.Index)
}
