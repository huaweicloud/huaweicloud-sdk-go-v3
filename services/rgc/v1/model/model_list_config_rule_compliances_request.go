package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListConfigRuleCompliancesRequest Request Object
type ListConfigRuleCompliancesRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`
}

func (o ListConfigRuleCompliancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConfigRuleCompliancesRequest struct{}"
	}

	return strings.Join([]string{"ListConfigRuleCompliancesRequest", string(data)}, " ")
}
