package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VoiceListResult 注册成功的声音列表，调用失败时无此字段。
type VoiceListResult struct {

	// 声色列表
	Voices *[]VoiceListResultVoices `json:"voices,omitempty"`
}

func (o VoiceListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceListResult struct{}"
	}

	return strings.Join([]string{"VoiceListResult", string(data)}, " ")
}
