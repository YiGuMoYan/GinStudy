package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// type User struct {
// 	Name string `json:"name" binding:"required,min=4,max=8"`
// 	Age  int    `json:"age" binding:"lt=30,gt=18"`
// 	Sex  string `json:"sex" binding:"oneof=男 女"`
// 	Data string `json:"data" binding:"datetime=2006-01-02 15:04:05"`
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.POST("/", func(context *gin.Context) {
// 		var user User
// 		err := context.ShouldBindJSON(&user)
// 		if err != nil {
// 			context.JSON(http.StatusOK, Response{http.StatusBadRequest, nil, err.Error()})
// 			return
// 		}
// 		context.JSON(http.StatusOK, Response{http.StatusOK, user, "成功"})
// 	})
//
// 	router.Run()
// }
