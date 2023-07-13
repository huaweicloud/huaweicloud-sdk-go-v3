package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubcustomerMonthlyBillsResponse Response Object
type ListSubcustomerMonthlyBillsResponse struct {

	// 账单记录，具体参考表2。
	BillSums *[]BillSumInfoV2 `json:"bill_sums,omitempty"`

	// 总记录数。
	Count *int32 `json:"count,omitempty"`

	// 货币编码。 CNY：人民币。
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubcustomerMonthlyBillsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubcustomerMonthlyBillsResponse struct{}"
	}

	return strings.Join([]string{"ListSubcustomerMonthlyBillsResponse", string(data)}, " ")
}
