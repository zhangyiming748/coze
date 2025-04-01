package main

import (
	"coze/voice"
	"log"
	"os"
)

func main() {
	// 检查参数数量
	if len(os.Args) < 3 {
		log.Fatalf("请输入语音ID和文本")
	}
	id := os.Args[1]
	text := os.Args[2]
	voice.GenerateAudio(id, text)
}
