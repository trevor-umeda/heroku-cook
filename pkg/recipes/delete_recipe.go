package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/trevor-umeda/heroku-cook/pkg/common/models"
	"net/http"
)

func (h handler) DeleteRecipe(c *gin.Context) {
	id := c.Param("id")

	var recipe models.Recipe
	if result := h.DB.First(&recipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&recipe)
	c.Status(http.StatusOK)
}
