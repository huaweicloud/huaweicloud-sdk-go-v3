package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListSubcustomerMonthlyBillsResponse struct {

	// 账单记录，具体参考表2。
	BillSums *[]BillSumInfoV2 `json:"bill_sums,omitempty" xml:"bill_sums"`

	// 总记录数。
	Count *int32 `json:"count,omitempty" xml:"count"`

	// 货币编码。 CNY：人民币。
	Currency       *string `json:"currency,omitempty" xml:"currency"`
	HttpStatusCode int     `json:"-"`
}

func (o ListSubcustomerMonthlyBillsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubcustomerMonthlyBillsResponse struct{}"
	}

	return strings.Join([]string{"ListSubcustomerMonthlyBillsResponse", string(data)}, " ")
}
