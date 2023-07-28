package main

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )
//
// func _string(context *gin.Context) {
// 	context.String(http.StatusOK, "你好")
// }
//
// func _json(context *gin.Context) {
// 	// json 响应结构体
// 	type User struct {
// 		UserName string
// 		Age      int
// 		Password string `json:"-"` // 忽略转化为 json
// 	}
// 	// user := User{"张三", 18, "123456"}
// 	// context.JSON(200, user)
//
// 	// json 响应map
// 	// userMap := map[string]string{
// 	// 	"username": "张三",
// 	// 	"age":      "23",
// 	// }
// 	// context.JSON(200, userMap)
//
// 	// 直接响应 json
// 	context.JSON(200, gin.H{"username": "张三", "age": 23})
//
// }
//
// func _html(context *gin.Context)  {
// 	context.HTML(http.StatusOK, "index.html", gin.H{"username":"张三"})
// }
//
// func _redirect(context *gin.Context) {
// 	context.Redirect(http.StatusMovedPermanently, "/html")
// }
//
// func main() {
// 	router := gin.Default()
//
// 	// 加载模版目录下所有文件
// 	router.LoadHTMLGlob("templates/*")
// 	// 在 golang 中，没有相对文件的路径，只有相对目录的路径
// 	router.StaticFile("static/Head.png", "static/Head.png")
//
// 	router.GET("/", _string)
// 	router.GET("/json", _json)
// 	router.GET("/html", _html)
// 	router.GET("/redirect", _redirect)
// 	router.Run(":8080")
//
// }
