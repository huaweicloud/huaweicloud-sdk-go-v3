package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAppRuleRequest Request Object
type DeleteAppRuleRequest struct {

	// 规则ID。
	RuleId string `json:"rule_id"`
}

func (o DeleteAppRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAppRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAppRuleRequest", string(data)}, " ")
}
