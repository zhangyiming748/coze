package util

import (
	"github.com/zhangyiming748/lumberjack"
	"io"
	"log"
	"os"
)

var logLocation string

func GetLog() string {
	return logLocation
}

func SetLog() {
	logLocation = "TTS.log"
	fileLogger := &lumberjack.Logger{
		Filename:   logLocation,
		MaxSize:    1, // MB
		MaxBackups: 1,
		MaxAge:     28, // days
	}
	fileLogger.Rotate()
	// 创建一个用于输出到控制台的Logger实例
	consoleLogger := log.New(os.Stdout, "CONSOLE: ", log.LstdFlags)

	// 设置文件Logger
	//log.SetOutput(fileLogger)

	// 同时输出到文件和控制台
	log.SetOutput(io.MultiWriter(fileLogger, consoleLogger.Writer()))
	log.SetFlags(log.Ltime | log.Lshortfile)
	// 在这里开始记录日志

	// 记录更多日志...

	// 关闭日志文件
	//defer fileLogger.Close()
}
