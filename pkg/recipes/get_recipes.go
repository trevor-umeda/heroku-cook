package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/trevor-umeda/heroku-cook/pkg/common/models"
	"net/http"
)

func (h handler) getRecipes(c *gin.Context) {
	var recipes []models.Recipe

	if result := h.DB.Find(&recipes); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	c.JSON(http.StatusOK, &recipes)
}
