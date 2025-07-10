package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListExternalConfigRuleCompliancesRequest Request Object
type ListExternalConfigRuleCompliancesRequest struct {

	// 纳管账号ID。
	ManagedAccountId string `json:"managed_account_id"`
}

func (o ListExternalConfigRuleCompliancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListExternalConfigRuleCompliancesRequest struct{}"
	}

	return strings.Join([]string{"ListExternalConfigRuleCompliancesRequest", string(data)}, " ")
}
