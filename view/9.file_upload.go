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
// 	router.POST("/upload", func(context *gin.Context) {
// 		file, _ := context.FormFile("file")
//
// 		fmt.Println(file.Filename)
// 		fmt.Println(file.Header)
// 		fmt.Println(file.Size / 1024)
// 		context.SaveUploadedFile(file, fmt.Sprintf("uploads/%s", file.Filename))
//
// 		// readerFile, _ := file.Open()
// 		// // data, _ := io.ReadAll(readerFile)
// 		// // fmt.Println(data)
// 		// writerFile, _ := os.Create(fmt.Sprintf("uploads/%s", file.Filename))
// 		// defer writerFile.Close()
// 		// n, _ := io.Copy(writerFile, readerFile)
// 		// fmt.Println(n)
//
// 		context.JSON(http.StatusOK, Response{http.StatusOK, nil, "上传成功"})
// 	})
//
// 	router.POST("/uploads", func(context *gin.Context) {
// 		form, _ := context.MultipartForm()
// 		files := form.File["upload[]"]
// 		for _, file := range files {
// 			context.SaveUploadedFile(file, fmt.Sprintf("uploads/%s", file.Filename))
// 		}
// 		context.JSON(http.StatusOK, Response{http.StatusOK, nil, fmt.Sprintf("上传成功 %d 个文件", len(files))})
// 	})
//
// 	router.Run()
// }
