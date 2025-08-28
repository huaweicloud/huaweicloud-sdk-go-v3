package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TtsJobFile struct {

	// 音频文件下载链接，有效期为1个小时。
	AudioFileUrl *string `json:"audio_file_url,omitempty"`

	// 时间戳文件下载链接，有效期为1个小时。
	AudioInfoFileUrl *string `json:"audio_info_file_url,omitempty"`

	// 字幕文件下载链接，有效期为1个小时。
	AudioSrtFileUrl *string `json:"audio_srt_file_url,omitempty"`

	// 动作分析文件下载链接，有效期为1个小时。
	AudioActionFileUrl *string `json:"audio_action_file_url,omitempty"`
}

func (o TtsJobFile) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtsJobFile struct{}"
	}

	return strings.Join([]string{"TtsJobFile", string(data)}, " ")
}
