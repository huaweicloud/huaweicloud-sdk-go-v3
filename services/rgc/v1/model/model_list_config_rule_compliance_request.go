package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigRuleComplianceRequest Request Object
type ListConfigRuleComplianceRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`
}

func (o ListConfigRuleComplianceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigRuleComplianceRequest struct{}"
	}

	return strings.Join([]string{"ListConfigRuleComplianceRequest", string(data)}, " ")
}
