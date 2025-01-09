package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductResult 产品计算结果
type ProductResult struct {

	// 对应请求体中的ID
	Id *string `json:"id,omitempty"`

	// 产品ID
	ProductId *string `json:"product_id,omitempty"`

	// 总额，即最终优惠后的金额(预留实例只包含预付部分)
	Amount float32 `json:"amount,omitempty"`

	// 优惠额(预留实例只包含预付部分)
	DiscountAmount float32 `json:"discount_amount,omitempty"`

	// 原总额，即优惠前总额(预留实例只包含预付部分)
	OriginalAmount float32 `json:"original_amount,omitempty"`

	// 官网价（非所有接口有该字段）(预留实例只包含预付部分)
	OfficialWebsiteAmount float32 `json:"official_website_amount,omitempty"`

	// 官网价优惠额（非所有接口有该字段）(预留实例只包含预付部分)
	OfficialWebsiteDiscountAmount float32 `json:"official_website_discount_amount,omitempty"`

	// 可选折扣优惠额(预留实例只包含预付部分)，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额（非所有接口有该字段）
	OptionalDiscountAmount float32 `json:"optional_discount_amount,omitempty"`

	// 总分期金额(批量询价的商品分期周期类型一致，才会有总分期金额，分期周期类型不一致，该信息没有)
	PerAmount float32 `json:"per_amount,omitempty"`

	// 总分期金额的优惠额(perDiscountAmount = perOriginalAmount - perAmount)
	PerDiscountAmount float32 `json:"per_discount_amount,omitempty"`

	// 总分期金额原价
	PerOriginalAmount float32 `json:"per_original_amount,omitempty"`

	// 分期金额的官网价
	PerOfficialWebsiteAmount float32 `json:"per_official_website_amount,omitempty"`

	// 分期金额的官网价官网价优惠额，即：perOfficialWebsiteDiscountAmount =perOriginalAmount-perOfficialWebsiteAmount
	PerOfficialWebsiteDiscountAmount float32 `json:"per_official_website_discount_amount,omitempty"`

	// 分期金额的可选折扣优惠额，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额 perOptionalDiscountAmount= perOfficialWebsiteAmount- perAmount
	PerOptionalDiscountAmount float32 `json:"per_optional_discount_amount,omitempty"`

	// 分期周期类型 2:月 4:小时
	PerPeriodType *int32 `json:"per_period_type,omitempty"`

	// 度量单位
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 扩展参数
	ExtendParams *string `json:"extend_params,omitempty"`
}

func (o ProductResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductResult struct{}"
	}

	return strings.Join([]string{"ProductResult", string(data)}, " ")
}
