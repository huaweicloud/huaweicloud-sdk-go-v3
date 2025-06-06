package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type OptionalDiscountRatingResult struct {

	// 折扣优惠ID。
	DiscountId *string `json:"discount_id,omitempty"`

	// 总额，即最终优惠后的金额。 amount= official_website_amount - discountAmount。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 包年/包月产品的官网价。
	OfficialWebsiteAmount *decimal.Decimal `json:"official_website_amount,omitempty"`

	// 可选折扣优惠额，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 价格度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 折扣优惠类型。商务授权折扣：605(华为云BE场景下的合同商务折扣)、伙伴授予折扣：606(分销商BE场景下的合同商务折扣)、伙伴设置折扣：607(合作伙伴设置折扣-折扣率)
	DiscountType *int32 `json:"discount_type,omitempty"`

	// 折扣名称。
	DiscountName *string `json:"discount_name,omitempty"`

	// 是否为最优折扣。 0：不是最优折扣，为缺省值。1：是最优折扣最优折扣：在商务折扣、伙伴折扣中选择（优惠金额最大的折扣为最优，优惠金额相等则按此顺序排优先级），促销折扣，折扣券不参与最优折扣的计算。
	BestOffer *int32 `json:"best_offer,omitempty"`

	// 产品询价结果，具体参见表5。
	ProductRatingResults *[]PeriodProductRatingResult `json:"product_rating_results,omitempty"`
}

func (o OptionalDiscountRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionalDiscountRatingResult struct{}"
	}

	return strings.Join([]string{"OptionalDiscountRatingResult", string(data)}, " ")
}
