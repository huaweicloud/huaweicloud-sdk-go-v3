package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleIdList **参数解释**： 规则ID列表
type RuleIdList struct {

	// 规则id列表
	Rules *[]RuleId `json:"rules,omitempty"`
}

func (o RuleIdList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleIdList struct{}"
	}

	return strings.Join([]string{"RuleIdList", string(data)}, " ")
}
