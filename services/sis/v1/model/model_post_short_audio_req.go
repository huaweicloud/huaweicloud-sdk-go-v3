package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostShortAudioReq struct {
	Config *Config `json:"config" xml:"config"`

	// 语音数据，base64编码，要求base64编码后大小不超过4M，音频时长不超过1分钟。
	Data string `json:"data" xml:"data"`
}

func (o PostShortAudioReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostShortAudioReq struct{}"
	}

	return strings.Join([]string{"PostShortAudioReq", string(data)}, " ")
}
