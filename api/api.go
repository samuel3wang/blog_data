package api

import (
	"fmt"
	"net/http"
	
	"blog_data/psql"
	"github.com/gin-gonic/gin"
)

func Getdata(r *gin.Context){
	var blog []Blog
	
	psql.DB.Find(&blog)

	


	r.JSON(http.StatusOK, gin.H{
		"message" : "all posts",
		"data" : blog,
	})
}
// ----------------------------------------------------------------
// func GetCategory(r *gin.Context){
// 	var blog []Blog

// 	psql.DB.Distinct("category").Find(&blog)

// 	r.JSON(http.StatusOK, gin.H{
// 		"message" : "all categories",
// 		"data" : blog,
// 	})

// }
func GetCategory(r *gin.Context){
	var categorys []Category

	psql.DB.Table("blogs").Distinct("category").Find(&categorys)

	r.JSON(http.StatusOK, gin.H{
		"message" : "all categories",
		"data" : categorys,
	})

}
// ----------------------------------------------------------------
func Show(r *gin.Context){
	blog := getByCategory(r)
	fmt.Println(blog)
	if blog.Title == "" {
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message" : "",
		"data" : blog,
	})
}

func getByCategory(r *gin.Context) Blog {
	var blog Blog
	category := r.Param("category")

	psql.DB.Find(&blog, "category = ?", category)
	
	if (blog.Title) == "" {
		r.JSON(http.StatusNotFound, gin.H{
			"message" : "can not find this person",
			"data" : "",
		})
	}
	return blog
}