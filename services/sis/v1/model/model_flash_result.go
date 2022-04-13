package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type FlashResult struct {
	// 音频声道id

	ChannelId *int32 `json:"channel_id,omitempty"`
	// 分句结果

	Sentences *[]Sentences `json:"sentences,omitempty"`
}

func (o FlashResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlashResult struct{}"
	}

	return strings.Join([]string{"FlashResult", string(data)}, " ")
}
