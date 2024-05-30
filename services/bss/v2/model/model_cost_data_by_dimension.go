package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CostDataByDimension struct {

	// 维度列表，具体请参见表 DimensionGroup。
	Dimensions *[]DimensionGroup `json:"dimensions,omitempty"`

	// 成本值，具体请参见表 Cost。
	Costs *[]Cost `json:"costs,omitempty"`

	// 此维度值对应整个时间跨度的成本汇总金额。
	AmountByCosts *string `json:"amount_by_costs,omitempty"`

	// 此维度值对应整个时间跨度的官网价汇总金额。
	OfficialAmountByCosts *string `json:"official_amount_by_costs,omitempty"`
}

func (o CostDataByDimension) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CostDataByDimension struct{}"
	}

	return strings.Join([]string{"CostDataByDimension", string(data)}, " ")
}
