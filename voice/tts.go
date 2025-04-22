package voice

import (
	"coze/util"
	"log"
	"strings"
)

const (
	TTS = "/audio/speech"
)

/*
curl -X POST 'https://api.coze.cn/v1/audio/speech' \
-H "Authorization: Bearer pat_i6d6nLon47ReVkEQOBT74wLoRpaCaP0ksj00SYZLrntGbKbEGqpmlouCF8esMtzO" \
-H "Content-Type: application/json" \
-d '{
  "input": "最近新买了一个 Mac Mini 完全覆盖了树莓派作为服务器的需求 把所有服务迁移到 Mac Mini 之后 树莓派终于闲置了 安装很久以前就宣称支持树莓派5的开放麒麟系统 看看效果如何",
  "voice_id": "7426720361733046281",
  "response_format": "mp3",
  "speed": 1.5,
  "sample_rate": 24000
}'
*/
func GenerateAudio(voice_id, text, speed string) {
	params := map[string]string{
		"input":           text,
		"voice_id":        voice_id,
		"response_format": "mp3",
		"speed":           speed,
		"sample_rate":     "24000",
	}
	hearers := map[string]string{
		"Authorization": "Bearer pat_i6d6nLon47ReVkEQOBT74wLoRpaCaP0ksj00SYZLrntGbKbEGqpmlouCF8esMtzO",
		"Content-Type":  "application/json",
	}
	host := "https://api.coze.cn/v1/audio/speech"
	fname := strings.Join([]string{text, "mp3"}, ".")
	if err := util.HttpPostJsoDownload(hearers, params, host, fname); err != nil {
		log.Println("GenerateAudio err:", err)
	}
}
