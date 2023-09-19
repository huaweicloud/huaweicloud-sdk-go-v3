package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShootScriptAudioFiles 用于语音驱动的音频文件上传URL列表。
type ShootScriptAudioFiles struct {

	// 用于语音驱动的音频文件上传URL。
	AudioFileUrl *[]ShootScriptAudioFileItem `json:"audio_file_url,omitempty"`
}

func (o ShootScriptAudioFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShootScriptAudioFiles struct{}"
	}

	return strings.Join([]string{"ShootScriptAudioFiles", string(data)}, " ")
}
