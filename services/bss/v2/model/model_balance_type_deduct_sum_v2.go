package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type BalanceTypeDeductSumV2 struct {

	// 账户类型。 BALANCE_TYPE_DEBIT：现金BALANCE_TYPE_CREDIT：信用BALANCE_TYPE_BONUS：奖励BALANCE_TYPE_COUPON：代金券BALANCE_TYPE_RCASH_COUPON 现金券。BALANCE_TYPE_STORED_VALUE_CARD：储值卡消费
	BalanceType *string `json:"balance_type,omitempty"`

	// 金额。 对于billType=1或者2的账单，该金额为负值。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 账单类型。 0：正常1：退订2：华为核销
	BillType *string `json:"bill_type,omitempty"`
}

func (o BalanceTypeDeductSumV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BalanceTypeDeductSumV2 struct{}"
	}

	return strings.Join([]string{"BalanceTypeDeductSumV2", string(data)}, " ")
}
