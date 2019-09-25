package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/thornshell/goblogV2/models"
	"net/http"
	"strconv"
)

func GetArticle(c *gin.Context) {
	i, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	article := models.GetArticle(i)
	c.JSON(http.StatusOK, article)
}
