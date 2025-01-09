package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OfficialWebsiteRatingResult 结果
type OfficialWebsiteRatingResult struct {

	// 订单总额，即最终优惠后的订单金额(预留实例只包含预付部分)
	Amount float32 `json:"amount,omitempty"`

	// 官网价(预留实例只包含预付部分)
	OfficialWebsiteAmount float32 `json:"official_website_amount,omitempty"`

	// 订单原总额，即优惠前订单总额(预留实例只包含预付部分)
	OriginalAmount float32 `json:"original_amount,omitempty"`

	// 官网价优惠额(预留实例只包含预付部分)
	OfficialWebsiteDiscountAmount float32 `json:"official_website_discount_amount,omitempty"`

	// 可选折扣优惠额(预留实例只包含预付部分)
	OptionalDiscountAmount float32 `json:"optional_discount_amount,omitempty"`

	// 总优惠额(预留实例只包含预付部分)
	DiscountAmount float32 `json:"discount_amount,omitempty"`

	// 总分期金额(批量询价的商品分期周期类型一致，才会有总分期金额，分期周期类型不一致，该信息没有)
	PerAmount float32 `json:"per_amount,omitempty"`

	// 总分期金额的优惠额(perDiscountAmount = perOriginalAmount - perAmount)
	PerDiscountAmount float32 `json:"per_discount_amount,omitempty"`

	// 总分期金额原价
	PerOriginalAmount float32 `json:"per_original_amount,omitempty"`

	// 分期周期类型 2:月 4:小时
	PerPeriodType *int32 `json:"per_period_type,omitempty"`

	// 度量单位
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 产品询价结果
	ProductRatingResults *[]ProductResult `json:"product_rating_results,omitempty"`
}

func (o OfficialWebsiteRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfficialWebsiteRatingResult struct{}"
	}

	return strings.Join([]string{"OfficialWebsiteRatingResult", string(data)}, " ")
}
