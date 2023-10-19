package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionInfo 过滤条件
type ConditionInfo struct {

	// 属性
	Attribute *string `json:"attribute,omitempty"`

	// 操作员，默认值EQ
	Operator *string `json:"operator,omitempty"`

	// 值
	Value *interface{} `json:"value,omitempty"`

	// 条件拼接准则，取值范围 AND,OR
	Condition *string `json:"condition,omitempty"`

	// 条件准则
	Criterion *[]DataMapFilterCriteria `json:"criterion,omitempty"`
}

func (o ConditionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionInfo struct{}"
	}

	return strings.Join([]string{"ConditionInfo", string(data)}, " ")
}
