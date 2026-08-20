package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRenewRateOnPeriodResponse Response Object
type ListRenewRateOnPeriodResponse struct {

	// 币种。USD：美元。值为空代表美元。
	Currency *string `json:"currency,omitempty"`

	// 主资源（包含从资源）询价结果。具体请参见表RenewInquiryResultInfo。
	RenewInquiryResults *[]RenewInquiryResultInfo `json:"renew_inquiry_results,omitempty"`

	OfficialWebsiteRatingResult *OfficialWebsiteRatingResultV2 `json:"official_website_rating_result,omitempty"`

	// 存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果，具体参见表OptionalDiscountRatingResultV2。
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResultV2 `json:"optional_discount_rating_results,omitempty"`

	// 失败的资源信息列表。具体请参见表FailResourceInfo。
	FailResourceInfos *[]FailResourceInfo `json:"fail_resource_infos,omitempty"`
	HttpStatusCode    int                 `json:"-"`
}

func (o ListRenewRateOnPeriodResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRenewRateOnPeriodResponse struct{}"
	}

	return strings.Join([]string{"ListRenewRateOnPeriodResponse", string(data)}, " ")
}
