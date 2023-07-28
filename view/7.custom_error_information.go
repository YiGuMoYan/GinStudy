package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// type User struct {
// 	Name string `json:"name" binding:"required" msg:"用户名校验失败"`
// 	Age  int    `json:"age" binding:"required" msg:"年龄校验失败"`
// }

// // GetValidMsg 获取结构体中 msg 参数
// func GetValidMsg(err error, obj any) string {
// 	getObj := reflect.TypeOf(obj)
// 	// 将 err 接口断言为具体类型
// 	if errs, ok := err.(validator.ValidationErrors); ok {
// 		// 断言成功
// 		for _, e := range errs {
// 			// 循环每一个错误信息
// 			// 根据报错字段，获取结构体具体字段
// 			if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
// 				return f.Tag.Get("msg")
// 			}
// 		}
// 	}
// 	return ""
// }
//
// func main() {
// 	router := gin.Default()
//
// 	router.POST("/", func(context *gin.Context) {
// 		var user User
// 		err := context.ShouldBindJSON(&user)
// 		if err != nil {
// 			context.JSON(http.StatusOK, Response{http.StatusBadRequest, nil, GetValidMsg(err, &user)})
// 			return
// 		}
// 		context.JSON(http.StatusOK, Response{http.StatusOK, user, "成功"})
// 	})
//
// 	router.Run()
// }
