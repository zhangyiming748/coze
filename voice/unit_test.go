package voice

import (
	"log"
	"os"
	"testing"
)
func init() {
	env:=os.Getenv("token")
	log.Println(env)
}
// go test -v -timeout 10h -run TestVoice
func TestVoiceList(t *testing.T) {
	GetVoiceList()
}

func TestGenerateAudio(t *testing.T) {
	id:="7426720361733046281"
	text:="我发现了一个暂时还是免费的文字转语音的api,叫'扣子'"
	GenerateAudio(id,text)
}
