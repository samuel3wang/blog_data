package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"blog_data/psql"
)
// query parameters
func RegisterRouter(){
	r := gin.Default()
	r.GET("/blogs", 		Getdata)
	r.GET("/blogs/:title", 	Show)

	// r.POST("/posts", 		Store)
	// r.PUT("/posts/:id", 	Update)
	// r.DELETE("/posts/:id", 	Delete)
	
	r.Run(":8088")
}
func Getdata(r *gin.Context){
	var blog []Blog
	
	psql.DB.Find(&blog)

	r.JSON(http.StatusOK, gin.H{
		"message" : "all posts",
		"data" : blog,
	})
}
func Show(r *gin.Context){
	blog := getById(r)
	if blog.Title == "" {
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message" : "",
		"data" : blog,
	})
}
// func Store(r *gin.Context){
// 	var person People
// 	if err := r.ShouldBindJSON(&person); err != nil {
// 		r.JSON(http.StatusBadRequest, gin.H{
// 			"message" : "the required element missed",
// 			"data" : "",
// 		})
// 	}
// 	psql.DB.Create(&person)
// 	r.JSON(http.StatusOK, gin.H{
// 		"message" : "all people",
// 		"data" : person,
// 	})
// }
// func Update(r *gin.Context){
// 	oldperson := getById(r)
// 	if oldperson.Name == ""{
// 		return
// 	}

// 	var update People
// 	if err := r.ShouldBindJSON(&update); err != nil {
// 		r.JSON(http.StatusBadRequest, gin.H{
// 			"message" : "the required element missed",
// 			"data" : "",
// 		})
// 	}

// 	oldperson.Ex = update.Ex
// 	oldperson.ID = update.ID
// 	oldperson.Name = update.Name
// 	psql.DB.Save(&oldperson)
// 	r.JSON(http.StatusOK, gin.H{
// 		"message" : "update successfully",
// 		"data" : oldperson,
// 	})
// }
// func Delete(r *gin.Context){
// 	var person People
// 	id := r.Param("id")
// 	psql.DB.First(&person, id)
// 	if person.Name == "" {
// 		r.JSON(http.StatusNotFound, gin.H{
// 			"message" : "can not find this person",
// 			"data" : "",
// 		})
// 		return
// 	}
	
// 	psql.DB.Delete(&person)
// 	r.JSON(http.StatusOK, gin.H{
// 		"message" : "",
// 		"data" : "delete successfully",
// 	})
// }
func getById(r *gin.Context) Blog{
	var blog Blog
	title := r.Param("title")
	psql.DB.First(&blog, title)
	if blog.Title == "" {
		r.JSON(http.StatusNotFound, gin.H{
			"message" : "can not find this person",
			"data" : "",
		})
	}
	return blog
}