package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Parser struct {

	// 数值
	ChannelReferCount *int64 `json:"channel_refer_count,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// UUID
	ParserId *string `json:"parser_id,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`
}

func (o Parser) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Parser struct{}"
	}

	return strings.Join([]string{"Parser", string(data)}, " ")
}
