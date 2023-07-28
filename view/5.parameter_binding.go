package main

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// 	Sex  string `json:"sex"`
// }
//
// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.POST("/", func(context *gin.Context) {
// 		var user User
// 		err := context.ShouldBindJSON(&user)
// 		if err != nil {
// 			context.JSON(http.StatusOK, Response{http.StatusBadRequest, nil, "参数错误"})
// 			return
// 		}
// 		context.JSON(http.StatusOK, Response{http.StatusOK, user, "成功"})
// 	})
//
// 	router.Run()
// }
