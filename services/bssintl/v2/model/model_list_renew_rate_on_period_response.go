package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRenewRateOnPeriodResponse Response Object
type ListRenewRateOnPeriodResponse struct {

	// |参数名称：币种。CNY：人民币。USD：美元。| |参数约束及描述：币种。CNY：人民币。USD：美元。|
	Currency *string `json:"currency,omitempty"`

	// |参数名称：主资源（包含从资源）询价结果| |参数约束以及描述：主资源（包含从资源）询价结果|
	RenewInquiryResults *[]RenewInquiryResultInfo `json:"renew_inquiry_results,omitempty"`

	OfficialWebsiteRatingResult *OfficialWebsiteRatingResultV2 `json:"official_website_rating_result,omitempty"`

	// |参数名称：存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果| |参数约束以及描述：存在可选折扣优惠时返回折扣优惠维度询价结果，每个折扣优惠一组询价结果|
	OptionalDiscountRatingResults *[]OptionalDiscountRatingResultV2 `json:"optional_discount_rating_results,omitempty"`

	// |参数名称：失败的资源信息列表| |参数约束以及描述：失败的资源信息列表|
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
