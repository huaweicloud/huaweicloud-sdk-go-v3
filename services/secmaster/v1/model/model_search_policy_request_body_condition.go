package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SearchPolicyRequestBodyCondition 查询条件
type SearchPolicyRequestBodyCondition struct {

	// 查询条件
	Conditions *[]SearchPolicyRequestBodyConditionConditions `json:"conditions,omitempty"`

	// 条件名称
	Logics *[]string `json:"logics,omitempty"`
}

func (o SearchPolicyRequestBodyCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SearchPolicyRequestBodyCondition struct{}"
	}

	return strings.Join([]string{"SearchPolicyRequestBodyCondition", string(data)}, " ")
}
