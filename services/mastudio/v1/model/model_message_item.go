package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MessageItem 多轮对话响应对象
type MessageItem struct {

	// 角色
	Role *string `json:"role,omitempty"`

	// 模型响应
	Content string `json:"content"`
}

func (o MessageItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessageItem struct{}"
	}

	return strings.Join([]string{"MessageItem", string(data)}, " ")
}
