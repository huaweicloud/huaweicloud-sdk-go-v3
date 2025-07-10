package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OptionalDiscountRatingResult 结果。
type OptionalDiscountRatingResult struct {

	// 折扣优惠Id。
	DiscountId *string `json:"discount_id,omitempty"`

	// 订单总额，即最终优惠后的订单金额(预留实例只包含预付部分)。
	Amount float32 `json:"amount,omitempty"`

	// 官网价(预留实例只包含预付部分)。
	OfficialWebsiteAmount float32 `json:"official_website_amount,omitempty"`

	// 订单原总额，即优惠前订单总额(预留实例只包含预付部分)。
	OriginalAmount float32 `json:"original_amount,omitempty"`

	// 官网价优惠额(预留实例只包含预付部分)。
	OfficialWebsiteDiscountAmount float32 `json:"official_website_discount_amount,omitempty"`

	// 可选折扣优惠额(预留实例只包含预付部分)。
	OptionalDiscountAmount float32 `json:"optional_discount_amount,omitempty"`

	// 总优惠额(预留实例只包含预付部分)。
	DiscountAmount float32 `json:"discount_amount,omitempty"`

	// 总分期金额(批量询价的商品分期周期类型一致，才会有总分期金额，分期周期类型不一致，该信息没有)。
	PerAmount float32 `json:"per_amount,omitempty"`

	// 总分期金额的优惠额(perDiscountAmount = perOriginalAmount - perAmount)。
	PerDiscountAmount float32 `json:"per_discount_amount,omitempty"`

	// 总分期金额原价。
	PerOriginalAmount float32 `json:"per_original_amount,omitempty"`

	// 总分期金额的官网价(批量询价的商品分期周期类型一致，才会有总分期金额的官网价，分期周期类型不一致，该信息没有)。
	PerOfficialWebsiteAmount float32 `json:"per_official_website_amount,omitempty"`

	// 总分期金额的官网价官网价优惠额，即： perOfficialWebsiteDiscountAmount =perOriginalAmount-perOfficialWebsiteAmount (批量询价的商品分期周期类型一致，才会有总分期金额的官网价，分期周期类型不一致，该信息没有)。
	PerOfficialWebsiteDiscountAmount float32 `json:"per_official_website_discount_amount,omitempty"`

	// 总分期金额的可选折扣优惠额，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额 perOptionalDiscountAmount= perOfficialWebsiteAmount- perAmount (批量询价的商品分期周期类型一致，才会有总分期金额的官网价，分期周期类型不一致，该信息没有)。
	PerOptionalDiscountAmount float32 `json:"per_optional_discount_amount,omitempty"`

	// 分期周期类型 2:月 4:小时。
	PerPeriodType *int32 `json:"per_period_type,omitempty"`

	// 度量单位。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 折扣优惠类型。
	DiscountType *int32 `json:"discount_type,omitempty"`

	// 折扣名称。
	DiscountName *string `json:"discount_name,omitempty"`

	// 是否为最优折扣0：不是最优折扣；为缺省值。1：是最优折扣；最优折扣：在商务折扣、伙伴折扣和促销折扣中选择（优惠金额相等则按此顺序排优先级），折扣券不参与最优折扣的计算。
	BestOffer *int32 `json:"best_offer,omitempty"`

	// sameRatioFlag。
	SameRatioFlag *int32 `json:"same_ratio_flag,omitempty"`

	// sameRatioFlag为1时有值，表示该折扣的折扣率。
	DiscountRatio float32 `json:"discount_ratio,omitempty"`

	// 折扣优惠基本信息；调用者在确定好折扣优惠后、下单时，使用此字段值，填入到订购/变更接口中的promotionInfo字段。
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 产品询价结果。
	ProductRatingResults *[]ProductResult `json:"product_rating_results,omitempty"`
}

func (o OptionalDiscountRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionalDiscountRatingResult struct{}"
	}

	return strings.Join([]string{"OptionalDiscountRatingResult", string(data)}, " ")
}
