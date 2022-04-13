package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 标签选择器
type LabelSelector struct {
	// 匹配规则表达式

	MatchExpressions *[]MatchExpression `json:"matchExpressions,omitempty"`
	// 匹配的标签，格式为key:value键值对。 单个键值对相当于matchExpressions的一个元素，key字段为key，操作符为In，values数组中只有value。

	MatchLabels map[string]string `json:"matchLabels,omitempty"`
}

func (o LabelSelector) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LabelSelector struct{}"
	}

	return strings.Join([]string{"LabelSelector", string(data)}, " ")
}
