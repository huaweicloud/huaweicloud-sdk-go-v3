package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPartnerBalancesResponse Response Object
type ListPartnerBalancesResponse struct {

	// 账户余额列表。 具体请参见表2
	AccountBalances *[]AccountBalanceV2 `json:"account_balances,omitempty"`
	HttpStatusCode  int                 `json:"-"`
}

func (o ListPartnerBalancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPartnerBalancesResponse struct{}"
	}

	return strings.Join([]string{"ListPartnerBalancesResponse", string(data)}, " ")
}
