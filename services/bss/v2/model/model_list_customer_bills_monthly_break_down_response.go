package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCustomerBillsMonthlyBreakDownResponse struct {
	// |参数名称：货币单位代码：CNY：人民币| |参数约束及描述：货币单位代码：CNY：人民币|

	Currency *string `json:"currency,omitempty"`
	// |参数名称：结果集数量，只有成功才返回这个参数。| |参数的约束及描述：结果集数量，只有成功才返回这个参数。|

	TotalCount *int32 `json:"total_count,omitempty"`
	// |参数名称：查询查询月度成本响应| |参数约束以及描述：查询查询月度成本响应|

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
