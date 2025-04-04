package main

import (
	"coze/util"
	"coze/voice"
	"os"
)

func main() {
	// 检查参数数量
	util.SetLog()
	if len(os.Args) < 4 {
		words:=util.ReadByLine("speech.txt")
		for _, word := range words {
			voice.GenerateAudio("7426720361733013513",word,"1.6")
		}
	}else if len(os.Args) == 4{
		id := os.Args[1]
		text := os.Args[2]
		speed := os.Args[3]
		voice.GenerateAudio(id, text,speed)
	}
}
