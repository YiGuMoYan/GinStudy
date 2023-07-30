package main

// type Response struct {
// 	Code int    `json:"code"`
// 	Data any    `json:"data"`
// 	Msg  string `json:"msg"`
// }
//
// func LogFormatterParams (params gin.LogFormatterParams) string {
// 	return fmt.Sprintf(
// 		"[GINSTUDY] %-30s|%s%-10d%s|%s%-10s%s|%-10s\n",
// 		params.TimeStamp.Format("2006-01-02 15:04:05"),
// 		params.StatusCodeColor(), params.StatusCode, params.ResetColor(),
// 		params.MethodColor(), params.Method, params.ResetColor(),
// 		params.Path,
// 	)
// }
//
// func main() {
// 	file, _ := os.Create("1.gin.log")
// 	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
// 	// router := gin.Default()
//
// 	router := gin.New()
//
// 	// router.Use(gin.LoggerWithFormatter(LogFormatterParams))
//
// 	router.Use(gin.LoggerWithConfig(gin.LoggerConfig{Formatter: LogFormatterParams}))
//
// 	router.GET("/", func(context *gin.Context) {
// 		context.JSON(http.StatusOK, Response{http.StatusOK, nil, "/"})
// 	})
//
// 	router.GET("/index", func(context *gin.Context) {
// 		context.JSON(http.StatusOK, Response{http.StatusOK, nil, "/index"})
// 	})
//
// 	router.Run(":8080")
// }
