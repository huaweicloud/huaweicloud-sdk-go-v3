/*
 * Bss
 *
 * Business Support System API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListRateOnPeriodDetailResponse struct {
	OfficialWebsiteRatingResult *OfficialWebsiteRatingResult `json:"official_website_rating_result,omitempty"`
	// |参数名称：存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果| |参数的约束及描述：存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果|
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResult `json:"optional_discount_rating_results,omitempty"`
	// |参数名称：币种| |参数约束及描述：比如CNY|
	Currency *string `json:"currency,omitempty"`
}

func (o ListRateOnPeriodDetailResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListRateOnPeriodDetailResponse", string(data)}, " ")
}
