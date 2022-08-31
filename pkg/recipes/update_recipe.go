package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/trevor-umeda/heroku-cook/pkg/common/models"
	"net/http"
)

type UpdateRecipeRequestBody struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
}

func (h handler) UpdateRecipe(c *gin.Context) {
	id := c.Param("id")

	body := UpdateRecipeRequestBody{}

	//getting requests body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var recipe models.Recipe

	if result := h.DB.First(&recipe, id); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	recipe.Title = body.Title
	recipe.Link = body.Link
	recipe.Description = body.Description
	recipe.Tags = body.Tags

	h.DB.Save(&recipe)

	c.JSON(http.StatusCreated, &recipe)
}
