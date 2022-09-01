package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMonthlyExpendituresResponse struct {

	// 返回码
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code"`

	// 返回码描述
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg"`

	// 货币单位代码 USD：美元
	Currency *string `json:"currency,omitempty" xml:"currency"`

	// 总条数
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 总金额（包含退订）。
	TotalAmount *float64 `json:"total_amount,omitempty" xml:"total_amount"`

	// 总欠费金额。
	DebtAmount *float64 `json:"debt_amount,omitempty" xml:"debt_amount"`

	// 代金券金额。
	CouponAmount *float64 `json:"coupon_amount,omitempty" xml:"coupon_amount"`

	// 现金券金额，预留。
	CashcouponAmount *float64 `json:"cashcoupon_amount,omitempty" xml:"cashcoupon_amount"`

	// 储值卡金额，预留。
	StoredcardAmount *float64 `json:"storedcard_amount,omitempty" xml:"storedcard_amount"`

	// 现金账户金额。
	DebitAmount *float64 `json:"debit_amount,omitempty" xml:"debit_amount"`

	// 信用账户金额。
	CreditAmount *float64 `json:"credit_amount,omitempty" xml:"credit_amount"`

	// 金额单位。 1：元3：分 默认值为3。
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 账单记录，具体参考表3。
	BillSums       *[]BillSumRecordInfo `json:"bill_sums,omitempty" xml:"bill_sums"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListMonthlyExpendituresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonthlyExpendituresResponse struct{}"
	}

	return strings.Join([]string{"ListMonthlyExpendituresResponse", string(data)}, " ")
}
