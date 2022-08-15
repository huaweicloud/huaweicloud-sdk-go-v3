package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListCostsReq struct {
	TimeCondition *TimeCondition `json:"time_condition"`

	// |参数名称：查询维度，具体请参见表 GroupBy。| |参数约束以及描述：查询维度，具体请参见表 GroupBy。|
	Groupby []GroupBy `json:"groupby"`

	// |参数名称：成本类型UNBLENDED_COST：原始成本AMORTIZED_COST：摊销成本| |参数约束及描述：成本类型UNBLENDED_COST：原始成本AMORTIZED_COST：摊销成本|
	CostType string `json:"cost_type"`

	// |参数名称：展示的金额类型PAYMENT_AMOUNT：应付NET_AMOUNT：实付| |参数约束及描述：展示的金额类型PAYMENT_AMOUNT：应付NET_AMOUNT：实付|
	AmountType string `json:"amount_type"`

	// |参数名称：偏移量| |参数的约束及描述：从0开始，默认为0|
	Offset *int32 `json:"offset,omitempty"`

	// |参数名称：每次查询的记录数| |参数的约束及描述：默认为10|
	Limit *int32 `json:"limit,omitempty"`

	// |参数名称：过滤条件，具体请参见表 过滤条件。| |参数约束以及描述：此参数不携带或携带值为空列表或携带值为null时，不作为筛选条件。|
	Filters *[]FilterV2 `json:"filters,omitempty"`
}

func (o ListCostsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCostsReq struct{}"
	}

	return strings.Join([]string{"ListCostsReq", string(data)}, " ")
}
