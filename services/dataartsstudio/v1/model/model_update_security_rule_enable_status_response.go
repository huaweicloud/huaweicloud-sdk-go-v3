package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityRuleEnableStatusResponse Response Object
type UpdateSecurityRuleEnableStatusResponse struct {

	// 识别规则id
	RuleId *string `json:"rule_id,omitempty"`

	// 识别规则名称
	RuleName *string `json:"rule_name,omitempty"`

	// 识别规则是否开启
	Enabled        *bool `json:"enabled,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o UpdateSecurityRuleEnableStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityRuleEnableStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdateSecurityRuleEnableStatusResponse", string(data)}, " ")
}
