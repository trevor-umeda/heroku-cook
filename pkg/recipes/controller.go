package recipes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := r.Group("/api/recipes")
	routes.POST("/", h.AddRecipe)
	routes.GET("/", h.getRecipes)
	routes.GET("/:id", h.getRecipe)
	routes.PUT("/:id", h.UpdateRecipe)
	routes.DELETE("/:id", h.DeleteRecipe)
}
