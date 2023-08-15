package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPostpaidBillSumResponse Response Object
type ListPostpaidBillSumResponse struct {

	// 账单所归属的月份。只有成功才返回这个参数。 格式：YYYY-MM
	BillCycle *string `json:"bill_cycle,omitempty"`

	// 账单中的应还金额（含税）。 应还金额（包含销项税）=消费金额+退款金额+调账金额
	InitialAmountDue *float64 `json:"initial_amount_due,omitempty"`

	// 账单中的消费金额。
	ConsumeAmount *float64 `json:"consume_amount,omitempty"`

	// 账单中的退款金额。
	Refunds *float64 `json:"refunds,omitempty"`

	// 账单中的调账金额，即伙伴在账期内的调账信息如：欠款核销金额等。
	Adjustments *float64 `json:"adjustments,omitempty"`

	// 账单中的销项税金额，销项税不计入应还金额。
	TaxAmount *float64 `json:"tax_amount,omitempty"`

	// 只有成功才返回这个参数。 美金：USD
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListPostpaidBillSumResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostpaidBillSumResponse struct{}"
	}

	return strings.Join([]string{"ListPostpaidBillSumResponse", string(data)}, " ")
}
