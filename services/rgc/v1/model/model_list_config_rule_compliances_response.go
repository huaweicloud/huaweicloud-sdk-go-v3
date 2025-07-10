package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigRuleCompliancesResponse Response Object
type ListConfigRuleCompliancesResponse struct {

	// 纳管账号ID
	AccountId *string `json:"account_id,omitempty"`

	// Config规则合规性信息
	ConfigRuleCompliances *[]ConfigRuleCompliance `json:"config_rule_compliances,omitempty"`
	HttpStatusCode        int                     `json:"-"`
}

func (o ListConfigRuleCompliancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigRuleCompliancesResponse struct{}"
	}

	return strings.Join([]string{"ListConfigRuleCompliancesResponse", string(data)}, " ")
}
