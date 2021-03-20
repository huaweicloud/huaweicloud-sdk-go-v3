package model

import (
	"encoding/json"

	"strings"
)

type TransferAmountInfoV2 struct {
	// 可拨款的金额。

	AvailTransferAmount *float64 `json:"avail_transfer_amount,omitempty"`
	// 金额单位。 1：元

	MeasureId *int32 `json:"measure_id,omitempty"`
	// 币种。 CNY：人民币

	Currency *string `json:"currency,omitempty"`
	// 账户余额（仅balance_type=信用账户时才有这个字段）。

	Amount *float64 `json:"amount,omitempty"`
	// 信用额度（仅balance_type=信用账户时才有这个字段）。

	CreditAmount *float64 `json:"credit_amount,omitempty"`
	// 信用额度过期时间。 UTC时间，格式为：2016-03-28T14:45:38Z。 （仅balance_type=信用账户时才有这个字段）。 如果查询信用账户可拨款余额的查询结果没有失效时间，表示永久有效。

	ExpireTime *string `json:"expire_time,omitempty"`
}

func (o TransferAmountInfoV2) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TransferAmountInfoV2 struct{}"
	}

	return strings.Join([]string{"TransferAmountInfoV2", string(data)}, " ")
}
