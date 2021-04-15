package v1

import (
	"goblogV2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetArticle(c *gin.Context) {
	i, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	article := models.GetArticle(i)
	c.JSON(http.StatusOK, article)
}
