package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RetrieveAmountInfoV2 struct {

	// 可回收的金额。
	AvailRetrieveAmount *float64 `json:"avail_retrieve_amount,omitempty" xml:"avail_retrieve_amount"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 币种。 CNY：人民币
	Currency *string `json:"currency,omitempty" xml:"currency"`

	// 账户余额（仅balance_type=信用账户时才有这个字段）。
	Amount *float64 `json:"amount,omitempty" xml:"amount"`

	// 信用额度（仅balance_type=信用账户时才有这个字段）。
	CreditAmount *float64 `json:"credit_amount,omitempty" xml:"credit_amount"`

	// 信用额度过期时间。 UTC时间，格式为：2016-03-28T14:45:38Z。 如果查询信用账户可回收余额的查询结果没有失效时间，表示永久有效。
	ExpireTime *string `json:"expire_time,omitempty" xml:"expire_time"`
}

func (o RetrieveAmountInfoV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetrieveAmountInfoV2 struct{}"
	}

	return strings.Join([]string{"RetrieveAmountInfoV2", string(data)}, " ")
}
