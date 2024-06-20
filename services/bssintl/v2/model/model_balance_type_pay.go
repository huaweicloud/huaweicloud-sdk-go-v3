package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type BalanceTypePay struct {

	// 账户类型。 BALANCE_TYPE_DEBIT：现金账户BALANCE_TYPE_CREDIT：信用账户BALANCE_TYPE_BONUS：奖励账户（该账户已下线）BALANCE_TYPE_COUPON：代金券账户BALANCE_TYPE_DEBIT_RATE：折扣账户
	BalanceTypeId *string `json:"balance_type_id,omitempty"`

	// 支出金额。 单位：分
	DeductAmount *decimal.Decimal `json:"deduct_amount,omitempty"`
}

func (o BalanceTypePay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BalanceTypePay struct{}"
	}

	return strings.Join([]string{"BalanceTypePay", string(data)}, " ")
}
