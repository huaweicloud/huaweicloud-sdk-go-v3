package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCustomerAccountBalancesResponse struct {

	// 账户余额列表。 具体请参见表1。
	AccountBalances *[]AccountBalanceV3 `json:"account_balances,omitempty" xml:"account_balances"`

	// 欠款总金额。
	DebtAmount *float64 `json:"debt_amount,omitempty" xml:"debt_amount"`

	// 度量单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 币种。 USD：美元。
	Currency       *string `json:"currency,omitempty" xml:"currency"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCustomerAccountBalancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerAccountBalancesResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerAccountBalancesResponse", string(data)}, " ")
}
