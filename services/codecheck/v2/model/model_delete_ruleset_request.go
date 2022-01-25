package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteRulesetRequest struct {
	// 规则集ID

	RulesetId string `json:"ruleset_id"`
}

func (o DeleteRulesetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRulesetRequest struct{}"
	}

	return strings.Join([]string{"DeleteRulesetRequest", string(data)}, " ")
}
