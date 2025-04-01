package main

import (
	"coze/util"
	"coze/voice"
	"log"
	"os"
)

func main() {
	// 检查参数数量
	util.SetLog()
	if len(os.Args) < 4 {
		log.Fatalf("请输入语音ID文本和语速")
	}
	id := os.Args[1]
	text := os.Args[2]
	speed := os.Args[3]
	voice.GenerateAudio(id, text,speed)
}
