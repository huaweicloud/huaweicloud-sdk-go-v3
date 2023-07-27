package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TextChoice 通用文本响应
type TextChoice struct {

	// 回复的索引
	Index int32 `json:"index"`

	// 模型响应
	Text string `json:"text"`
}

func (o TextChoice) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextChoice struct{}"
	}

	return strings.Join([]string{"TextChoice", string(data)}, " ")
}
