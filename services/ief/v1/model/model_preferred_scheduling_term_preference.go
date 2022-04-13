package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 与相应权重关联的节点选择器项。
type PreferredSchedulingTermPreference struct {
	// 匹配规则表达式

	MatchExpressions *[]MatchExpression `json:"matchExpressions,omitempty"`
}

func (o PreferredSchedulingTermPreference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PreferredSchedulingTermPreference struct{}"
	}

	return strings.Join([]string{"PreferredSchedulingTermPreference", string(data)}, " ")
}
