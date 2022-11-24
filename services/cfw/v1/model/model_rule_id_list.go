package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 规则id列表
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
