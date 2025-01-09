package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EstimateDesktopPoolExtendVolumeResponse Response Object
type EstimateDesktopPoolExtendVolumeResponse struct {

	// 币种，比如CNY
	Currency *string `json:"currency,omitempty"`

	// 扩展参数
	ExtendParams *string `json:"extend_params,omitempty"`

	OfficialWebsiteRatingResult *OfficialWebsiteRatingResult `json:"official_website_rating_result,omitempty"`

	// 存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResult `json:"optional_discount_rating_results,omitempty"`
	HttpStatusCode                int                             `json:"-"`
}

func (o EstimateDesktopPoolExtendVolumeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EstimateDesktopPoolExtendVolumeResponse struct{}"
	}

	return strings.Join([]string{"EstimateDesktopPoolExtendVolumeResponse", string(data)}, " ")
}
