package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LivePlayingShootScriptItem 直播话术配置。
type LivePlayingShootScriptItem struct {

	// 剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	// 段落标题。
	Title *string `json:"title,omitempty"`
}

func (o LivePlayingShootScriptItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LivePlayingShootScriptItem struct{}"
	}

	return strings.Join([]string{"LivePlayingShootScriptItem", string(data)}, " ")
}
