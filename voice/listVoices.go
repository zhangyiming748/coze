package voice

import (
	"coze/constant"
	"coze/util"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	LIST = "/audio/voices"
)

type Voice struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Detail struct {
		Logid string `json:"logid"`
	} `json:"detail"`
	Data struct {
		HasMore   bool `json:"has_more"`
		VoiceList []struct {
			CreateTime             int    `json:"create_time"`
			UpdateTime             int    `json:"update_time"`
			LanguageName           string `json:"language_name"`
			LanguageCode           string `json:"language_code"`
			AvailableTrainingTimes int    `json:"available_training_times"`
			IsSystemVoice          bool   `json:"is_system_voice"`
			Name                   string `json:"name"`
			PreviewText            string `json:"preview_text"`
			PreviewAudio           string `json:"preview_audio"`
			VoiceId                string `json:"voice_id"`
		} `json:"voice_list"`
	} `json:"data"`
}

// 在 Voice struct 后添加新的函数
func DownloadPreviewAudio(audioURL, fileName string) error {
	// 创建下载目录
	downloadDir := "downloads/voices"
	err := os.MkdirAll(downloadDir, 0755)
	if err != nil {
		return fmt.Errorf("创建下载目录失败: %v", err)
	}
	filePath := filepath.Join(downloadDir, fileName)
	// 构建 wget 命令
	cmd := exec.Command("wget", "--header", fmt.Sprintf("Authorization: %s", os.Getenv("TOKEN")), "-O", filePath, audioURL)
	// 执行命令
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("下载失败: %v, 输出: %s", err, string(output))
	}
	log.Printf("音频文件已保存至: %s\n", filePath)
	util.Sleep(3)
	return nil
}

func GetVoiceList() {
	token := os.Getenv("TOKEN")
	headers := map[string]string{
		"Authorization": token,
		"Content-Type":  "application/json",
	}
	date := map[string]string{
		"filter_system_voice": "false",
		"page_num":            "1",
		"page_size":           "200",
	}
	host := strings.Join([]string{constant.HOST, LIST}, "/")
	out, err := util.HttpGet(headers, date, host)
	if err != nil {
		log.Fatalf("http get 发生错误: %v\n", err)
	}
	log.Println(string(out))
	var voice Voice
	err = json.Unmarshal(out, &voice)
	if err != nil {
		log.Fatalf("解析json发生错误%v\n", err)
	}
	if voice.Code != 0 {
		log.Fatalf("请求服务器发生500错误,返回:%+v\n", voice)
	}
	for _, v := range voice.Data.VoiceList {
		log.Printf("创建时间:%v\n", formatTimestamp(v.CreateTime))
		log.Printf("更新时间:%v\n", formatTimestamp(v.UpdateTime))
		log.Printf("语言:%v\n", v.LanguageName)
		log.Printf("语言代码:%v\n", v.LanguageCode)
		log.Printf("训练次数:%v\n", v.AvailableTrainingTimes)
		log.Printf("是否为系统声音:%v\n", v.IsSystemVoice)
		log.Printf("音色名称:%v\n", v.Name)
		log.Printf("试听文本:%v\n", v.PreviewText)
		log.Printf("试听音频:%v\n", v.PreviewAudio)
		// 下载音频文件
		if v.PreviewAudio != "" {
			// 构建文件名
			fileName := fmt.Sprintf("%v. %s.mp3", v.VoiceId,v.Name)
			if err := DownloadPreviewAudio(v.PreviewAudio, fileName); err != nil {
				log.Printf("下载音频失败: %v\n", err)
			}
		}
		log.Printf("音色id:%v\n", v.VoiceId)
	}
}

func formatTimestamp(timestamp int) string {
    // 创建东八区时区对象
    cst := time.FixedZone("CST", 8*3600)
    // 转换时间戳并设置时区
    return time.Unix(int64(timestamp), 0).In(cst).Format("2006-01-02 15:04:05")
}