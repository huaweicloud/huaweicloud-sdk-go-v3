package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// ProductQuotingItemStep 产品报价项阶梯信息
type ProductQuotingItemStep struct {

	// 阶梯ID
	StepId *string `json:"step_id,omitempty"`

	// 阶梯编号
	StepNo *string `json:"step_no,omitempty"`

	// 阶梯起始值
	StepStart *decimal.Decimal `json:"step_start,omitempty"`

	// 起始值度量单位
	StartMeasureId *int32 `json:"start_measure_id,omitempty"`

	// 阶梯结束值
	StepEnd *decimal.Decimal `json:"step_end,omitempty"`

	// 结束值度量单位（1：元/美元）
	EndMeasureId *int32 `json:"end_measure_id,omitempty"`

	// 优惠方式：0：产品折扣，1：固定单价
	PreferentialType *int32 `json:"preferential_type,omitempty"`

	// 固定单价（preferential_type=1固定单价时有值）
	SalesPrice *decimal.Decimal `json:"sales_price,omitempty"`

	// 折扣率（preferential_type=0产品折扣时有值）
	DiscountRatio *decimal.Decimal `json:"discount_ratio,omitempty"`

	// 计费单位
	PricingBasis *string `json:"pricing_basis,omitempty"`
}

func (o ProductQuotingItemStep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductQuotingItemStep struct{}"
	}

	return strings.Join([]string{"ProductQuotingItemStep", string(data)}, " ")
}
