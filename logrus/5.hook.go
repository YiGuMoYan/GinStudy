package main

// type MyHook struct {
// }
//
// func (m MyHook) Levels() []logrus.Level {
// 	return []logrus.Level{logrus.ErrorLevel}
// }
//
// func (m MyHook) Fire(entry *logrus.Entry) error {
// 	fmt.Println(entry)
//
// 	file, _ := os.OpenFile("logrus/error.log", os.O_CREATE|os.O_CREATE|os.O_APPEND, 0666)
// 	line, _ := entry.String()
// 	file.Write([]byte(line))
//
// 	return nil
// }
//
// func main() {
// 	logrus.AddHook(&MyHook{})
//
// 	logrus.Errorf("你好")
// 	logrus.Infof("你好")
// }
