package main

import (
	"blog_data/psql"
	"blog_data/routers"
)
func main()  {

	psql.ConnectDB()

	r := routers.RegisterRouter()
		
	r.Run(":8088")

}