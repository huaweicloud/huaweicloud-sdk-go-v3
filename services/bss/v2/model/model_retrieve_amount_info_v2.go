package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type RetrieveAmountInfoV2 struct {

	// 可回收的金额。
	AvailRetrieveAmount *decimal.Decimal `json:"avail_retrieve_amount,omitempty"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 币种。 CNY：人民币
	Currency *string `json:"currency,omitempty"`

	// 账户余额（仅balance_type=信用账户时这个字段才有值）。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 信用额度（仅balance_type=信用账户时这个字段才有值）。
	CreditAmount *decimal.Decimal `json:"credit_amount,omitempty"`

	// 信用额度过期时间。 UTC时间，格式为：2016-03-28T14:45:38Z。 如果查询信用账户可回收余额的查询结果没有失效时间，表示永久有效。
	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o RetrieveAmountInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrieveAmountInfoV2 struct{}"
	}

	return strings.Join([]string{"RetrieveAmountInfoV2", string(data)}, " ")
}
