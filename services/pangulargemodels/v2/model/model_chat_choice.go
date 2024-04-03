package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChatChoice 通用文本响应
type ChatChoice struct {

	// 回复的索引
	Index int32 `json:"index"`

	Message *MessageItem `json:"message"`
}

func (o ChatChoice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatChoice struct{}"
	}

	return strings.Join([]string{"ChatChoice", string(data)}, " ")
}
