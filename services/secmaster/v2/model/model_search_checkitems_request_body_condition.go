package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchCheckitemsRequestBodyCondition 搜索条件表达式
type SearchCheckitemsRequestBodyCondition struct {

	// 表达式列表
	Conditions *[]BaselineSearchRequestBodyConditionConditions `json:"conditions,omitempty"`

	// 表达式名称列表
	Logics *[]string `json:"logics,omitempty"`
}

func (o SearchCheckitemsRequestBodyCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchCheckitemsRequestBodyCondition struct{}"
	}

	return strings.Join([]string{"SearchCheckitemsRequestBodyCondition", string(data)}, " ")
}
