package recipes

import (
	"github.com/gin-gonic/gin"
	"github.com/trevor-umeda/heroku-cook/pkg/common/models"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type AddRecipeRequestBody struct {
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
	Tags        string `json:"tags"`
	Rating      int    `json:"rating"`
}

func (h handler) AddRecipe(c *gin.Context) {
	body := AddRecipeRequestBody{}

	//getting requests body
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var recipe models.Recipe

	recipe.Title = body.Title
	recipe.Link = body.Link
	recipe.Description = body.Description
	recipe.Tags = body.Tags
	recipe.Rating = body.Rating
	if result := h.DB.Create(&recipe); result.Error != nil {
		c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}
	h.AddTagsInternal(recipe.ID, recipe.Tags)
	c.JSON(http.StatusCreated, &recipe)
}

func (h handler) AddTagsInternal(recipeId uint, t string) {
	var tag models.Tag

	tags := strings.Split(t, ",")
	for _, s := range tags {
		log.Default().Println("Adding tag " + s)
		result := h.DB.Where("tag = ?", s).First(&tag)

		// New tag
		if result.Error != nil {
			var newTag models.Tag
			newTag.Tag = s
			newTag.Recipes = strconv.FormatUint(uint64(recipeId), 10)
			if result := h.DB.Create(&newTag); result.Error != nil {
				return
			}
		} else {
			recipeList := strings.Split(tag.Recipes, ",")
			recipeList = append(recipeList, strconv.FormatUint(uint64(recipeId), 10))

			h.DB.Model(&tag).Update("recipes", strings.Join(recipeList, ","))
		}

	}
}
