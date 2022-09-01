package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListOnDemandResourceRatingsResponse struct {

	// 折扣的金额。
	Amount *float64 `json:"amount,omitempty" xml:"amount"`

	// 优惠额（官网价和总价的差）。
	DiscountAmount *float64 `json:"discount_amount,omitempty" xml:"discount_amount"`

	// 按需产品的官网价。
	OfficialWebsiteAmount *float64 `json:"official_website_amount,omitempty" xml:"official_website_amount"`

	// 度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 币种。 CNY：人民币。 值为空代表人民币。
	Currency *string `json:"currency,omitempty" xml:"currency"`

	// 产品询价结果，具体参见表2。
	ProductRatingResults *[]DemandProductRatingResult `json:"product_rating_results,omitempty" xml:"product_rating_results"`
	HttpStatusCode       int                          `json:"-"`
}

func (o ListOnDemandResourceRatingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListOnDemandResourceRatingsResponse struct{}"
	}

	return strings.Join([]string{"ListOnDemandResourceRatingsResponse", string(data)}, " ")
}
