package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// func m10(context *gin.Context)  {
// 	context.Set("name", "张三")
// 	fmt.Println("m10 in")
// }
//
// func m11(context *gin.Context)  {
// 	fmt.Println("m11 in")
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.Use(m10, m11)
//
// 	router.GET("/", func(context *gin.Context) {
// 		fmt.Println(context.Get("name"))
// 		context.JSON(http.StatusOK, Response{http.StatusOK, nil, "index"})
// 	})
//
// 	router.Run()
// }
