package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPostpaidBillSumRequest Request Object
type ListPostpaidBillSumRequest struct {

	// 账单所归属的月份。格式：YYYY-MM。
	BillCycle string `json:"bill_cycle"`
}

func (o ListPostpaidBillSumRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPostpaidBillSumRequest struct{}"
	}

	return strings.Join([]string{"ListPostpaidBillSumRequest", string(data)}, " ")
}
