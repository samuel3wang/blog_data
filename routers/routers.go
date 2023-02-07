package routers

import (
	"blog_data/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRouter() *gin.Engine{
	r := gin.Default()
	
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	r.Use(cors.New(corsConfig))

	r.GET("/blogs",								api.Getdata)
	r.GET("/blog/:title_id",			api.WholePost)
	r.GET("/category",						api.GetCategory)	
	r.GET("/category/:category",	api.Show)

	return r
}