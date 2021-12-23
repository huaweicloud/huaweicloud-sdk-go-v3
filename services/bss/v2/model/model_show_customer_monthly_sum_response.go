package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCustomerMonthlySumResponse struct {
	// 总条数，必须大于等于0。

	TotalCount *int32 `json:"total_count,omitempty"`
	// 账单记录，具体参考表2。

	BillSums *[]BillSumRecordInfoV2 `json:"bill_sums,omitempty"`
	// 总金额（包含退订）。

	ConsumeAmount *float64 `json:"consume_amount,omitempty"`
	// 总欠费金额。

	DebtAmount *float64 `json:"debt_amount,omitempty"`
	// 代金券金额。

	CouponAmount *float64 `json:"coupon_amount,omitempty"`
	// 现金券金额，预留。

	FlexipurchaseCouponAmount *float64 `json:"flexipurchase_coupon_amount,omitempty"`
	// 储值卡金额，预留。

	StoredValueCardAmount *float64 `json:"stored_value_card_amount,omitempty"`
	// 现金账户金额。

	CashAmount *float64 `json:"cash_amount,omitempty"`
	// 信用账户金额。

	CreditAmount *float64 `json:"credit_amount,omitempty"`
	// 欠费核销金额。

	WriteoffAmount *float64 `json:"writeoff_amount,omitempty"`
	// 金额单位。 1：元

	MeasureId *int32 `json:"measure_id,omitempty"`
	// 币种。 CNY：人民币。

	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCustomerMonthlySumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerMonthlySumResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerMonthlySumResponse", string(data)}, " ")
}
