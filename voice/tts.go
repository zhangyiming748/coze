package voice

import (
	"coze/constant"
	"coze/util"
	"log"
	"os"
	"strings"
)

const (
	TTS = "/audio/speech"
)

func GenerateAudio(voice_id,text string) {
	token := os.Getenv("TOKEN")
	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}
	date := map[string]string{
		"input":           text,
		"voice_id":        voice_id,
		"response_format": "mp3",
		"speed":           "1",
		"sample_rate":     "24000",
	}
	host := strings.Join([]string{constant.HOST, TTS}, "/")
	fname:=strings.Join([]string{text,"mp3"},".")
	err:=util.HttpPostJsoDownload(headers,date,host,fname)
	if err!= nil {
		log.Fatalf("生成音频失败: %v\n", err)
	}
}
