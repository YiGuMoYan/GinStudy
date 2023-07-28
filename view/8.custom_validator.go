package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// type User struct {
// 	Name string `json:"name" binding:"required,sign" msg:"用户名校验失败"`
// 	Age  int    `json:"age" binding:"required" msg:"年龄校验失败"`
// }
//
// func signValidator(fl validator.FieldLevel) bool {
// 	nameList := []string{"张三", "李四", "王五"}
// 	for _, nameStr := range nameList{
// 		name := fl.Field().Interface().(string)
// 		if name == nameStr {
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func main() {
// 	router := gin.Default()
//
// 	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 		v.RegisterValidation("sign", signValidator)
// 	}
//
// 	router.POST("/", func(context *gin.Context) {
// 		var user User
// 		err := context.ShouldBindJSON(&user)
//
// 		if err != nil {
// 			context.JSON(http.StatusOK, Response{http.StatusBadRequest, nil, err.Error()})
// 			return
// 		}
//
// 		context.JSON(http.StatusOK, Response{http.StatusOK, user, "成功"})
// 	})
//
// 	router.Run()
// }
