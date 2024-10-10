package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WafCustomRuleAction struct {

	// 防护动作。 “block”：拦截。 “pass”：放行。 “log”：仅记录
	Category *string `json:"category,omitempty"`
}

func (o WafCustomRuleAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WafCustomRuleAction struct{}"
	}

	return strings.Join([]string{"WafCustomRuleAction", string(data)}, " ")
}
