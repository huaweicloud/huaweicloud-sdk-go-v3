package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionItem 条件定义
type ConditionItem struct {

	// 条件名称
	Name *string `json:"name,omitempty"`

	// 条件详情
	Detail *string `json:"detail,omitempty"`

	// 条件表达式数据
	Data *[]string `json:"data,omitempty"`
}

func (o ConditionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionItem struct{}"
	}

	return strings.Join([]string{"ConditionItem", string(data)}, " ")
}
