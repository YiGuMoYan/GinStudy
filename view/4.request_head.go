package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.GET("/", func(context *gin.Context) {
// 		fmt.Println(context.GetHeader("jwt"))
// 		if context.GetHeader("jwt") == "" {
// 			context.JSON(http.StatusOK, Response{http.StatusBadRequest, "", "缺少jwt"})
// 			return
// 		}
// 		context.JSON(http.StatusOK, Response{http.StatusOK, "", "成功"})
// 	})
//
// 	router.Run(":8080")
// }
