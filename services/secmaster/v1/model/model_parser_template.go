package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ParserTemplate struct {

	// 描述信息
	Description *string `json:"description,omitempty"`

	// UUID
	ParserId *string `json:"parser_id,omitempty"`

	// 相关标题
	Title *string `json:"title,omitempty"`
}

func (o ParserTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ParserTemplate struct{}"
	}

	return strings.Join([]string{"ParserTemplate", string(data)}, " ")
}
