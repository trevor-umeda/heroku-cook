package tags

import (
	"github.com/gin-gonic/gin"
	"github.com/trevor-umeda/heroku-cook/pkg/common/models"
	"net/http"
	"strings"
)

func (h handler) getByTags(c *gin.Context) {
	p := c.Param("tag")
	var tag models.Tag

	if result := h.DB.Where("tag = ?", p).First(&tag); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	var recipes []models.Recipe
	ids := strings.Split(tag.Recipes, ",")

	if result := h.DB.Find(&recipes, ids); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	c.JSON(http.StatusOK, &recipes)
}
