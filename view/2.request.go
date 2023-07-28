package main

// import (
// 	"fmt"
// 	"github.com/gin-gonic/gin"
// )
//
// func _query(context *gin.Context) {
// 	// user := context.Query("user")
// 	// fmt.Println(user)
//
// 	user, ok := context.GetQuery("user")
// 	if !ok {
// 		fmt.Println("未传入字符串")
// 		return
// 	}
// 	fmt.Println(user)
// }
//
// func _param(content *gin.Context) {
// 	fmt.Println(content.Param("user_id"))
// 	fmt.Println(content.Param("book_id"))
// }
//
// func _form(content *gin.Context) {
// 	fmt.Println(content.PostForm("name"))
// 	// 如果用户不传，就是用默认值
// 	fmt.Println(content.DefaultPostForm("id", "12345"))
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/query", _query)
// 	router.GET("/param/:user_id/:book_id", _param)
// 	router.POST("/form", _form)
// 	router.Run()
// }
