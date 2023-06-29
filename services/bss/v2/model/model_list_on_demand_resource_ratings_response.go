package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

// ListOnDemandResourceRatingsResponse Response Object
type ListOnDemandResourceRatingsResponse struct {

	// 折扣的金额。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 优惠额（官网价和总价的差）。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 按需产品的官网价。
	OfficialWebsiteAmount *decimal.Decimal `json:"official_website_amount,omitempty"`

	// 度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 币种。 CNY：人民币。 值为空代表人民币。
	Currency *string `json:"currency,omitempty"`

	// 产品询价结果，具体参见表2。
	ProductRatingResults *[]DemandProductRatingResult `json:"product_rating_results,omitempty"`
	HttpStatusCode       int                          `json:"-"`
}

func (o ListOnDemandResourceRatingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOnDemandResourceRatingsResponse struct{}"
	}

	return strings.Join([]string{"ListOnDemandResourceRatingsResponse", string(data)}, " ")
}
