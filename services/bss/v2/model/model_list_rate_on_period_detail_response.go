package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRateOnPeriodDetailResponse Response Object
type ListRateOnPeriodDetailResponse struct {
	OfficialWebsiteRatingResult *OfficialWebsiteRatingResult `json:"official_website_rating_result,omitempty"`

	// 存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果，具体参见表4。
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResult `json:"optional_discount_rating_results,omitempty"`

	// 币种。 CNY：人民币。 值为空代表人民币。
	Currency       *string `json:"currency,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRateOnPeriodDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRateOnPeriodDetailResponse struct{}"
	}

	return strings.Join([]string{"ListRateOnPeriodDetailResponse", string(data)}, " ")
}
