package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCustomerAccountBalancesRequest Request Object
type ShowCustomerAccountBalancesRequest struct {
}

func (o ShowCustomerAccountBalancesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerAccountBalancesRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerAccountBalancesRequest", string(data)}, " ")
}
