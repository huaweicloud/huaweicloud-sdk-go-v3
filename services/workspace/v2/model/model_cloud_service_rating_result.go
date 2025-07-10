package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudServiceRatingResult 询价结果。
type CloudServiceRatingResult struct {

	// 下单请求体中的orderRequestId。
	OrderRequestId *string `json:"order_request_id,omitempty"`

	OfficialWebsiteRatingResult *OfficialWebsiteRatingResult `json:"official_website_rating_result,omitempty"`

	// 优惠询价结果。
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResult `json:"optional_discount_rating_results,omitempty"`
}

func (o CloudServiceRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudServiceRatingResult struct{}"
	}

	return strings.Join([]string{"CloudServiceRatingResult", string(data)}, " ")
}
