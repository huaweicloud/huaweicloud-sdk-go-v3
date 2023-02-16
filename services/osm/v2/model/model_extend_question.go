package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExtendQuestion struct {

	// 扩展问内容
	Content *string `json:"content,omitempty"`

	// 分值
	Score *float64 `json:"score,omitempty"`
}

func (o ExtendQuestion) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtendQuestion struct{}"
	}

	return strings.Join([]string{"ExtendQuestion", string(data)}, " ")
}
