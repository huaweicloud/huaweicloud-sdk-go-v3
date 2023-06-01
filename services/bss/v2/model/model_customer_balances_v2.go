package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type CustomerBalancesV2 struct {

	// 客户账号ID。
	CustomerId string `json:"customer_id"`

	// 客户欠款总额度。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 客户可用总额度。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 币种。 CNY：人民币。
	Currency *string `json:"currency,omitempty"`

	// 度量单位： 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o CustomerBalancesV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerBalancesV2 struct{}"
	}

	return strings.Join([]string{"CustomerBalancesV2", string(data)}, " ")
}
