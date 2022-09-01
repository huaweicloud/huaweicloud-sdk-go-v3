package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type FlashResult struct {

	// 音频声道id
	ChannelId *int32 `json:"channel_id,omitempty" xml:"channel_id"`

	// 分句结果
	Sentences *[]Sentences `json:"sentences,omitempty" xml:"sentences"`
}

func (o FlashResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlashResult struct{}"
	}

	return strings.Join([]string{"FlashResult", string(data)}, " ")
}
