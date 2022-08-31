package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/trevor-umeda/heroku-cook/pkg/common/db"
	"github.com/trevor-umeda/heroku-cook/pkg/recipes"
	"github.com/trevor-umeda/heroku-cook/pkg/tags"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DATABASE_URL").(string)

	r := gin.Default()
	h := db.Init(dbUrl)

	recipes.RegisterRoutes(r, h)
	tags.RegisterRoutes(r, h)

	r.Run(port)
}
