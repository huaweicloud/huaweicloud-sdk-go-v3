package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalConfigRuleCompliancesResponse Response Object
type ListExternalConfigRuleCompliancesResponse struct {

	// 纳管账号ID
	AccountId *string `json:"account_id,omitempty"`

	// Config规则合规性信息
	ConfigRuleCompliances *[]ExternalConfigRuleCompliance `json:"config_rule_compliances,omitempty"`
	HttpStatusCode        int                             `json:"-"`
}

func (o ListExternalConfigRuleCompliancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalConfigRuleCompliancesResponse struct{}"
	}

	return strings.Join([]string{"ListExternalConfigRuleCompliancesResponse", string(data)}, " ")
}
