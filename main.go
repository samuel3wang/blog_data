package main

import (
	"github.com/gin-gonic/gin"
	"blog_data/psql"
	"blog_data/api"
)
func main()  {
	gin.SetMode(gin.ReleaseMode)

	psql.ConnectDB()
	api.RegisterRouter()
}

// https://www.youtube.com/watch?v=alcanyAiv14&list=PLYp_Kd32Xvcq2Qyxu8pnJ0ZGAz9ewBZ3H&index=8