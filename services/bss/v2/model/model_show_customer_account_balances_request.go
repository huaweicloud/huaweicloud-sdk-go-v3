package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCustomerAccountBalancesRequest struct {
}

func (o ShowCustomerAccountBalancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerAccountBalancesRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerAccountBalancesRequest", string(data)}, " ")
}
