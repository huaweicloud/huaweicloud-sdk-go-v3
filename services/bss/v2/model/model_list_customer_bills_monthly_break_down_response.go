package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerBillsMonthlyBreakDownResponse Response Object
type ListCustomerBillsMonthlyBreakDownResponse struct {

	// 货币单位代码： CNY：人民币
	Currency *string `json:"currency,omitempty"`

	// 结果集数量，只有成功才返回这个参数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 分摊成本记录数据。 具体请参见表3。
	Details        *[]NvlCostAnalysedBillDetail `json:"details,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListCustomerBillsMonthlyBreakDownResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerBillsMonthlyBreakDownResponse struct{}"
	}

	return strings.Join([]string{"ListCustomerBillsMonthlyBreakDownResponse", string(data)}, " ")
}
