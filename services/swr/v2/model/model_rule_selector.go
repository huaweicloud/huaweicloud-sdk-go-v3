package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSelector struct {

	// 匹配类型，目前只支持doublestar
	Kind string `json:"kind"`

	// 选择器匹配类型,当前支持repoMatches
	Decoration string `json:"decoration"`

	// 选择器匹配样式，最大长度512。支持正则表达式，正则表达式规则可填写如 nginx-* ，{repo1，repo2} 等，其中： * ： 匹配不包含 '/' 的任何字段。 **：匹配包含 '/' 的任何字段。 ? ：匹配任何单个非 '/' 的字符。 {选项1,选项2,...}：同时匹配多个选项。
	Pattern string `json:"pattern"`

	// 预留字段
	Extras *string `json:"extras,omitempty"`
}

func (o RuleSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSelector struct{}"
	}

	return strings.Join([]string{"RuleSelector", string(data)}, " ")
}
