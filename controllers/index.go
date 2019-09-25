package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(ctx *gin.Context) {
	name := ctx.DefaultQuery("name", "a wei")
	ctx.String(http.StatusOK, "Hello %s", name)

}
