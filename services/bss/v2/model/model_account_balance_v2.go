package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AccountBalanceV2 struct {

	// 账户标识。
	AccountId string `json:"account_id"`

	// 账户类型： 1：余额2：信用5：奖励7：保证金8：可拨款
	AccountType int32 `json:"account_type"`

	// 账户余额。
	Amount float64 `json:"amount"`

	// 币种。 CNY：人民币。
	Currency string `json:"currency"`

	// 专款专用余额。
	DesignatedAmount *float64 `json:"designated_amount,omitempty"`

	// 总信用额度。只有account_type取值为2时，该字段才有效。
	CreditAmount *float64 `json:"credit_amount,omitempty"`

	// 度量单位。 1：元。
	MeasureId int32 `json:"measure_id"`

	// 备注。
	Memo *string `json:"memo,omitempty"`
}

func (o AccountBalanceV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccountBalanceV2 struct{}"
	}

	return strings.Join([]string{"AccountBalanceV2", string(data)}, " ")
}
