package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// ShowCustomerMonthlySumResponse Response Object
type ShowCustomerMonthlySumResponse struct {

	// 总条数，必须大于等于0。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 账单记录，具体参考表2。
	BillSums *[]BillSumRecordInfoV2 `json:"bill_sums,omitempty"`

	// 总金额（包含退订）。
	ConsumeAmount *decimal.Decimal `json:"consume_amount,omitempty"`

	// 总欠费金额。
	DebtAmount *decimal.Decimal `json:"debt_amount,omitempty"`

	// 代金券金额。
	CouponAmount *decimal.Decimal `json:"coupon_amount,omitempty"`

	// 现金券金额，预留。
	FlexipurchaseCouponAmount *decimal.Decimal `json:"flexipurchase_coupon_amount,omitempty"`

	// 储值卡金额，预留。
	StoredValueCardAmount *decimal.Decimal `json:"stored_value_card_amount,omitempty"`

	// 现金账户金额。
	CashAmount *decimal.Decimal `json:"cash_amount,omitempty"`

	// 信用账户金额。
	CreditAmount *decimal.Decimal `json:"credit_amount,omitempty"`

	// 欠费核销金额。
	WriteoffAmount *decimal.Decimal `json:"writeoff_amount,omitempty"`

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
