package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OptionalDiscountRatingResultV2 struct {

	// |参数名称：折扣优惠ID| |参数约束及描述：折扣优惠ID|
	DiscountId *string `json:"discount_id,omitempty"`

	// |参数名称：总额，即最终优惠后的金额。单位为元。amount= official_website_amount - discount_amount| |参数约束及描述：总额，即最终优惠后的金额。单位为元。amount= official_website_amount - discount_amount|
	Amount *string `json:"amount,omitempty"`

	// |参数名称：官网价。单位为元| |参数约束及描述：官网价。单位为元|
	OfficialWebsiteAmount *string `json:"official_website_amount,omitempty"`

	// |参数名称：可选折扣优惠额，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额。单位为| |参数约束及描述：可选折扣优惠额，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额。单位为|
	DiscountAmount *string `json:"discount_amount,omitempty"`

	// |参数名称：折扣优惠类型。合同商务折扣：605：华为云BE场景下的合同商务折扣606：分销商BE场景下的合同商务折扣伙伴授予折扣：607：合作伙伴授予折扣-折扣率| |参数的约束及描述：折扣优惠类型。合同商务折扣：605：华为云BE场景下的合同商务折扣606：分销商BE场景下的合同商务折扣伙伴授予折扣：607：合作伙伴授予折扣-折扣率|
	DiscountType *int32 `json:"discount_type,omitempty"`

	// |参数名称：折扣名称| |参数约束及描述：折扣名称|
	DiscountName *string `json:"discount_name,omitempty"`

	// |参数名称：是否为最优折扣。0：不是最优折扣，为缺省值。1：是最优折扣最优折扣：在商务折扣、伙伴折扣中选择（优惠金额最大的折扣为最优，优惠金额相等则按此顺序排优先级），促销折扣，折扣券不参与最优折扣的计算| |参数的约束及描述：是否为最优折扣。0：不是最优折扣，为缺省值。1：是最优折扣最优折扣：在商务折扣、伙伴折扣中选择（优惠金额最大的折扣为最优，优惠金额相等则按此顺序排优先级），促销折扣，折扣券不参与最优折扣的计算|
	BestOffer *int32 `json:"best_offer,omitempty"`

	// |参数名称：分期金额的官网价。单位为元| |参数约束及描述：分期金额的官网价。单位为元|
	InstallmentOfficialWebsiteAmount *string `json:"installment_official_website_amount,omitempty"`

	// |参数名称：分期金额的折扣价。单位为元| |参数约束及描述：分期金额的折扣价。单位为元|
	InstallmentOfficialDiscountAmount *string `json:"installment_official_discount_amount,omitempty"`

	// |参数名称：分期金额的成交价。单位为元| |参数约束及描述：分期金额的成交价。单位为元|
	InstallmentAmount *string `json:"installment_amount,omitempty"`

	// |参数名称：分期付款的周期类型。2：月| |参数的约束及描述：分期付款的周期类型。2：月|
	InstallmentPeriodType *int32 `json:"installment_period_type,omitempty"`
}

func (o OptionalDiscountRatingResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OptionalDiscountRatingResultV2 struct{}"
	}

	return strings.Join([]string{"OptionalDiscountRatingResultV2", string(data)}, " ")
}
