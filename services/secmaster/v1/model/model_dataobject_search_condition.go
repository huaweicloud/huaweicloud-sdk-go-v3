package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataobjectSearchCondition 搜索条件表达式
type DataobjectSearchCondition struct {

	// 表达式列表
	Conditions *[]DataobjectSearchConditionConditions `json:"conditions,omitempty"`

	// 表达式名称列表
	Logics *[]string `json:"logics,omitempty"`
}

func (o DataobjectSearchCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataobjectSearchCondition struct{}"
	}

	return strings.Join([]string{"DataobjectSearchCondition", string(data)}, " ")
}
