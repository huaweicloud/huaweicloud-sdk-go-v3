package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PeriodBatchUpChangeResourceRsp 桌面池升配变更询价响应体。
type PeriodBatchUpChangeResourceRsp struct {

	// 币种，比如CNY。
	Currency *string `json:"currency,omitempty"`

	// 扩展参数。
	ExtendParams *string `json:"extend_params,omitempty"`

	OfficialWebsiteRatingResult *OfficialWebsiteRatingResult `json:"official_website_rating_result,omitempty"`

	// 存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果。
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResult `json:"optional_discount_rating_results,omitempty"`
}

func (o PeriodBatchUpChangeResourceRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodBatchUpChangeResourceRsp struct{}"
	}

	return strings.Join([]string{"PeriodBatchUpChangeResourceRsp", string(data)}, " ")
}
