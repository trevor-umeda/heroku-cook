package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trevor-umeda/heroku-cook/pkg/common/db"
	"github.com/trevor-umeda/heroku-cook/pkg/recipes"
	"github.com/trevor-umeda/heroku-cook/pkg/tags"
	"net/http"
	"os"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := os.Getenv("PORT")
	if port == "" {
		port = viper.Get("PORT").(string)
	}
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		dbUrl = viper.Get("DATABASE_URL").(string)
	}

	r := gin.Default()
	h := db.Init(dbUrl)

	recipes.RegisterRoutes(r, h)
	tags.RegisterRoutes(r, h)
	r.LoadHTMLFiles("static/index.html")
	r.Static("/assets", "static/assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.Run(":" + port)
}
