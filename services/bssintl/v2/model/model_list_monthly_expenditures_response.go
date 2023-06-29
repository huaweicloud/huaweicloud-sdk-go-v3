package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMonthlyExpendituresResponse Response Object
type ListMonthlyExpendituresResponse struct {

	// 返回码
	ErrorCode *string `json:"error_code,omitempty"`

	// 返回码描述
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 货币单位代码 USD：美元
	Currency *string `json:"currency,omitempty"`

	// 总条数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 总金额（包含退订）。
	TotalAmount *float64 `json:"total_amount,omitempty"`

	// 总欠费金额。
	DebtAmount *float64 `json:"debt_amount,omitempty"`

	// 代金券金额。
	CouponAmount *float64 `json:"coupon_amount,omitempty"`

	// 现金券金额，预留。
	CashcouponAmount *float64 `json:"cashcoupon_amount,omitempty"`

	// 储值卡金额，预留。
	StoredcardAmount *float64 `json:"storedcard_amount,omitempty"`

	// 现金账户金额。
	DebitAmount *float64 `json:"debit_amount,omitempty"`

	// 信用账户金额。
	CreditAmount *float64 `json:"credit_amount,omitempty"`

	// 金额单位。 1：元3：分 默认值为3。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 账单记录，具体参考表3。
	BillSums       *[]BillSumRecordInfo `json:"bill_sums,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ListMonthlyExpendituresResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMonthlyExpendituresResponse struct{}"
	}

	return strings.Join([]string{"ListMonthlyExpendituresResponse", string(data)}, " ")
}
