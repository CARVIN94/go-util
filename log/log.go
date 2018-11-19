package log

import (
	"log"

	"github.com/CARVIN94/go-util/logcolor"
)

// Print 输出一切
func Print(msg interface{}) {
	log.Print(msg)
}

// Success 输出成功
func Success(msg string) {
	log.Print(logcolor.Success(msg))
}

// Warning 输出警告
func Warning(msg string) {
	log.Print(logcolor.Warning(msg))
}

// Info 输出信息
func Info(msg string) {
	log.Print(logcolor.Info(msg))
}

// Error 输出错误
func Error(err error, msg string) {
	log.Print(logcolor.Error(msg), err)
}

// Connect 连接输出
func Connect(mold string, status string, url string) {
	log.Print(logcolor.Magenta("["+mold+"] ") + status + ": " + url)
}

// Fatal 输出错误并关闭程序
func Fatal(msg string) {
	log.Fatal(logcolor.Error(msg))
}

// FailOnError 当错误时输出并关闭程序
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatal(logcolor.Error(msg), err.Error())
	}
}
