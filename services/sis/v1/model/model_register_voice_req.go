package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RegisterVoiceReq struct {
	Config *RegisterVoiceReqConfig `json:"config"`

	// 录音数据，使用base64编码，大小不超过6MB。支持wav、mp3、m4a格式，采样率不小于16kHz，时长在5-25秒，支持单、双通道。
	Data string `json:"data"`
}

func (o RegisterVoiceReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterVoiceReq struct{}"
	}

	return strings.Join([]string{"RegisterVoiceReq", string(data)}, " ")
}
