package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type PeriodProductRatingResult struct {

	// ID标识，来源于请求中的ID。
	Id *string `json:"id,omitempty"`

	// 包年/包月产品的ID。
	ProductId *string `json:"product_id,omitempty"`

	// 总额，即最终优惠后的金额。 amount= official_website_amount - discountAmount。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 包年/包月产品的官网价。
	OfficialWebsiteAmount *decimal.Decimal `json:"official_website_amount,omitempty"`

	// 可选折扣优惠额，如商务折扣、伙伴折扣、促销折扣和折扣券选用时的优惠额。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 价格度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o PeriodProductRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodProductRatingResult struct{}"
	}

	return strings.Join([]string{"PeriodProductRatingResult", string(data)}, " ")
}
