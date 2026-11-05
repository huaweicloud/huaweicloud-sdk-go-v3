package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AiguardCustomRegexBody struct {

	// 正则名称
	Name string `json:"name"`

	// 正则表达式
	RegexExpression string `json:"regex_expression"`
}

func (o AiguardCustomRegexBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiguardCustomRegexBody struct{}"
	}

	return strings.Join([]string{"AiguardCustomRegexBody", string(data)}, " ")
}
