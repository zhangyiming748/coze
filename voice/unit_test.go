package voice

import (
	"coze/util"
	"testing"
)

func init() {
	util.SetLog()
}

// go test -v -timeout 10h -run TestVoice
func TestVoiceList(t *testing.T) {
	GetVoiceList("1")
	GetVoiceList("2")
}

// go test -v -timeout 10h -run TestGenerateAudio
func TestGenerateAudio(t *testing.T) {
	id := "7426720361733046281"
	//text:="我发现了一个暂时还是免费的文字转语音的api,叫'扣子'"
	texts := []string{"最近新买了一个 Mac Mini", "完全覆盖了树莓派作为服务器的需求", "把所有服务迁移到 Mac Mini 之后", "树莓派终于闲置了", "安装很久以前就宣称支持树莓派5的开放麒麟系统", "看看效果如何"}
	//GenerateAudio(id,text,"1.1")
	for _, text := range texts {
		GenerateAudio(id, text, "1.5")
	}
}
