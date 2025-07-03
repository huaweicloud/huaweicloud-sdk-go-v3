package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectFilterV2 struct {

	// 操作符 in/like/startwith/endwith/=/!=/>/<等
	Operator string `json:"operator"`

	// 字段名称
	Field string `json:"field"`

	// 分组
	Group *string `json:"group,omitempty"`

	// 条件名称
	Name *string `json:"name,omitempty"`

	// 匹配方式
	MatchType *string `json:"match_type,omitempty"`

	// 优先级处理方式
	PriorityType *string `json:"priority_type,omitempty"`

	// 搜索值
	Values []interface{} `json:"values"`
}

func (o ObjectFilterV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectFilterV2 struct{}"
	}

	return strings.Join([]string{"ObjectFilterV2", string(data)}, " ")
}
