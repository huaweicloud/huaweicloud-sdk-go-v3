package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// ShowCustomerAccountBalancesResponse Response Object
type ShowCustomerAccountBalancesResponse struct {

	// 账户余额列表。 具体请参见表1。
	AccountBalances *[]AccountBalanceV3 `json:"account_balances,omitempty"`

	// 欠款总金额。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 度量单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 币种。 CNY：人民币。
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCustomerAccountBalancesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerAccountBalancesResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerAccountBalancesResponse", string(data)}, " ")
}
