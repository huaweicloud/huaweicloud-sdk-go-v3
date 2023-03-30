package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OfficialWebsiteRatingResultV2 struct {

	// |参数名称：官网价格。单位为元| |参数约束及描述：官网价格。单位为元|
	OfficialWebsiteAmount *string `json:"official_website_amount,omitempty"`

	// |参数名称：分期金额的官网价。单位为元| |参数约束及描述：分期金额的官网价。单位为元|
	InstallmentOfficialWebsiteAmount *string `json:"installment_official_website_amount,omitempty"`

	// |参数名称：分期付款的周期类型。2：月| |参数的约束及描述：分期付款的周期类型。2：月|
	InstallmentPeriodType *int32 `json:"installment_period_type,omitempty"`
}

func (o OfficialWebsiteRatingResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfficialWebsiteRatingResultV2 struct{}"
	}

	return strings.Join([]string{"OfficialWebsiteRatingResultV2", string(data)}, " ")
}
