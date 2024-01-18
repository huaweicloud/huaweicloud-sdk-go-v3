package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigRuleComplianceResponse Response Object
type ListConfigRuleComplianceResponse struct {

	// 纳管账号ID
	AccountId *string `json:"account_id,omitempty"`

	// Config规则合规性信息
	ConfigRuleCompliances *[]ConfigRuleCompliance `json:"config_rule_compliances,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ListConfigRuleComplianceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigRuleComplianceResponse struct{}"
	}

	return strings.Join([]string{"ListConfigRuleComplianceResponse", string(data)}, " ")
}
