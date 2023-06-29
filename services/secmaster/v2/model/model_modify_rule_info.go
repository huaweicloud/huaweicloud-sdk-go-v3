package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyRuleInfo Information of rule
type ModifyRuleInfo struct {
	Rule *ConditionInfo `json:"rule,omitempty"`
}

func (o ModifyRuleInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyRuleInfo struct{}"
	}

	return strings.Join([]string{"ModifyRuleInfo", string(data)}, " ")
}
