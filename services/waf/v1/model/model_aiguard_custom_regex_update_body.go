package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomRegexUpdateBody struct {

	// 正则名称
	Name *string `json:"name,omitempty"`

	// 正则表达式
	RegexExpression *string `json:"regex_expression,omitempty"`
}

func (o AiguardCustomRegexUpdateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomRegexUpdateBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomRegexUpdateBody", string(data)}, " ")
}
