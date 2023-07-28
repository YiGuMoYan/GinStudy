package main

// type ArticleModel struct {
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }
//
// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// func _getList(context *gin.Context) {
// 	articleList := []ArticleModel{
// 		{"Go", "GoContent"},
// 		{"Python", "PythonContent"},
// 		{"Java", "JavaContent"},
// 	}
// 	context.JSON(http.StatusOK, Response{http.StatusOK, articleList, "成功"})
// }
//
// func _getDetail(context *gin.Context) {
// 	fmt.Println(context.Param("id"))
// 	article := ArticleModel{"Go", "GoContent"}
//
// 	context.JSON(http.StatusOK, Response{http.StatusOK, article, "成功"})
// }
//
// func _create(context *gin.Context) {
// 	var article ArticleModel
// 	err := context.BindJSON(&article)
// 	if err != nil {
// 		context.JSON(http.StatusOK, Response{http.StatusBadRequest, nil, err.Error()})
// 		return
// 	}
// 	fmt.Println(article)
// 	context.JSON(http.StatusOK, Response{http.StatusOK, nil, "成功"})
// }
//
// func _update(context *gin.Context) {
// 	fmt.Println(context.Param("id"))
// 	var article ArticleModel
// 	err := context.BindJSON(&article)
// 	if err != nil {
// 		context.JSON(http.StatusOK, Response{http.StatusBadRequest, nil, err.Error()})
// 		return
// 	}
// 	fmt.Println(article)
// 	context.JSON(http.StatusOK, Response{http.StatusOK, nil, "成功"})
// }
//
// func _delete(context *gin.Context) {
// 	fmt.Println(context.Param("id"))
// 	context.JSON(http.StatusOK, Response{http.StatusOK, nil, "成功"})
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.GET("/articles", _getList)
// 	router.GET("/articles/:id", _getDetail)
// 	router.POST("/articles", _create)
// 	router.PUT("/articles/:id", _update)
// 	router.DELETE("/articles/:id", _delete)
//
// 	router.Run(":8080")
// }
