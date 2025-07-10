package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Content 控制策略内容。
type Content struct {

	// 英文策略内容。
	En *string `json:"en,omitempty"`

	// 中文策略内容。
	Ch *string `json:"ch,omitempty"`
}

func (o Content) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Content struct{}"
	}

	return strings.Join([]string{"Content", string(data)}, " ")
}
