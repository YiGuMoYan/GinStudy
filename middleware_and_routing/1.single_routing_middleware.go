package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// func m1(context *gin.Context)  {
// 	fmt.Println("m1...")
// 	context.JSON(http.StatusOK, Response{http.StatusOK, nil, "m1"})
// 	context.Abort()
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.GET("/", m1, func(context *gin.Context) {
// 		fmt.Println("index...")
// 		context.JSON(http.StatusOK, Response{http.StatusOK, nil, "index"})
// 	})
//
// 	router.Run()
// }
