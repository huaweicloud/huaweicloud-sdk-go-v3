package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowCustomerMonthlySumResponse struct {

	// 总条数，必须大于等于0。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 账单记录，具体参考表2。
	BillSums *[]BillSumRecordInfoV2 `json:"bill_sums,omitempty" xml:"bill_sums"`

	// 总金额（包含退订）。
	ConsumeAmount *float64 `json:"consume_amount,omitempty" xml:"consume_amount"`

	// 总欠费金额。
	DebtAmount *float64 `json:"debt_amount,omitempty" xml:"debt_amount"`

	// 代金券金额。
	CouponAmount *float64 `json:"coupon_amount,omitempty" xml:"coupon_amount"`

	// 现金券金额，预留。
	FlexipurchaseCouponAmount *float64 `json:"flexipurchase_coupon_amount,omitempty" xml:"flexipurchase_coupon_amount"`

	// 储值卡金额，预留。
	StoredValueCardAmount *float64 `json:"stored_value_card_amount,omitempty" xml:"stored_value_card_amount"`

	// 现金账户金额。
	CashAmount *float64 `json:"cash_amount,omitempty" xml:"cash_amount"`

	// 信用账户金额。
	CreditAmount *float64 `json:"credit_amount,omitempty" xml:"credit_amount"`

	// 欠费核销金额。
	WriteoffAmount *float64 `json:"writeoff_amount,omitempty" xml:"writeoff_amount"`

	// 金额单位。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 币种。 CNY：人民币。
	Currency       *string `json:"currency,omitempty" xml:"currency"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowCustomerMonthlySumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerMonthlySumResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomerMonthlySumResponse", string(data)}, " ")
}
