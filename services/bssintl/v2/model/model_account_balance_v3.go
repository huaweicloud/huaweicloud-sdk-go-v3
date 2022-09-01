package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountBalanceV3 struct {

	// 账户标识。
	AccountId string `json:"account_id" xml:"account_id"`

	// 账户类型。 1：余额2：信用5：奖励金7：保证金
	AccountType int32 `json:"account_type" xml:"account_type"`

	// 账户余额。
	Amount float64 `json:"amount" xml:"amount"`

	// 币种。 USD：美元。
	Currency string `json:"currency" xml:"currency"`

	// 专款专用余额。
	DesignatedAmount *float64 `json:"designated_amount,omitempty" xml:"designated_amount"`

	// 总信用额度，仅信用账户存在该字段。
	CreditAmount *float64 `json:"credit_amount,omitempty" xml:"credit_amount"`

	// 度量单位。 1：元
	MeasureId int32 `json:"measure_id" xml:"measure_id"`
}

func (o AccountBalanceV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountBalanceV3 struct{}"
	}

	return strings.Join([]string{"AccountBalanceV3", string(data)}, " ")
}
