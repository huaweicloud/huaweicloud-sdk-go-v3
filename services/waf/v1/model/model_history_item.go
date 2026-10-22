package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HistoryItem 对话历史单条消息
type HistoryItem struct {

	// 消息角色（user/assistant/system）
	Role *string `json:"role,omitempty"`

	// 该轮消息内容
	Content *string `json:"content,omitempty"`
}

func (o HistoryItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HistoryItem struct{}"
	}

	return strings.Join([]string{"HistoryItem", string(data)}, " ")
}
