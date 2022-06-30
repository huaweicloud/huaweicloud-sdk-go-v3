package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 匹配规则表达式
type MatchExpression struct {

	// 规则的标签
	Key *string `json:"key,omitempty"`

	// 操作符，取值如下： - In：标签值需要在values的列表中 - NotIn：标签的值不在某个列表中 - Exists：某个标签存在 - DoesNotExist：某个标签不存在 - Gt：标签的值大于某个值（字符串比较） - Lt：标签的值小于某个值（字符串比较）
	Operator *string `json:"operator,omitempty"`

	// 一组标签值。 - 如果运算符为In或NotIn，则值数组必须非空。 - 如果运算符为Exists 或DoesNotExist，则值数组必须为空。 - 如果运算符是Gt或Lt，则值数组必须具有单个元素，该元素将被解释为整数。
	Values *[]string `json:"values,omitempty"`
}

func (o MatchExpression) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MatchExpression struct{}"
	}

	return strings.Join([]string{"MatchExpression", string(data)}, " ")
}
