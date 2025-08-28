package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomerExpenditureQuotaRequest Request Object
type ShowCustomerExpenditureQuotaRequest struct {
}

func (o ShowCustomerExpenditureQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerExpenditureQuotaRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerExpenditureQuotaRequest", string(data)}, " ")
}
