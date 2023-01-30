package api

import (
	"fmt"
	"net/http"
	"strings"

	"blog_data/psql"
	"github.com/gin-gonic/gin"
)

// 設定字數
func limitString(str string, limit int) string {
	str = strings.TrimSpace(strings.Replace(str, "\n", "", -1))
	return string([]rune(str)[:limit])
}

// query home 預覽
func Getdata(r *gin.Context){
	var blog []Blog
	
	psql.DB.Find(&blog)

		var i int
		for i = 0 ; i<len(blog) ; i++{
			shortText := limitString(blog[i].Content, 95) + " ..."
			blog[i].Content = shortText
		}

	r.JSON(http.StatusOK, gin.H{
		"message" : "all posts",
		"data" : blog,
	})
}
// ----------------------------------------------------------------

// get category list
func GetCategory(r *gin.Context){
	var categorys []Category

	psql.DB.Table("blogs").Distinct("category").Find(&categorys)

	r.JSON(http.StatusOK, gin.H{
		"message" : "all categories",
		"data" : categorys,
	})

}
// ----------------------------------------------------------------
// query category title
func Show(r *gin.Context)  {
	var title []struct{
		Title_id int `json:"title_id"`
		Title string `json:"title"`
	}
	category := r.Param("category")
	psql.DB.Model(&Blog{}).Select("title_id", "title").Where("category = ?", category).Find(&title)

	r.JSON(http.StatusOK, gin.H{
		"message" : "",
		"data" : title,
	})
}
// ----------------------------------------------------------------
// query whole post
func WholePost(r *gin.Context)  {
	var blog Blog

	title_id := r.Param("title_id")
	psql.DB.Where("title_id = ?", title_id).Find(&blog)

	r.JSON(http.StatusOK, gin.H{
		"message" : "",
		"data" : blog,
	})
}
// -------------------------------------------------------------------
func Fmt(){
	fmt.Println("aaa")
}