package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OfficialWebsiteRatingResultV2 struct {

	// 官网价格。单位为美元
	OfficialWebsiteAmount *string `json:"official_website_amount,omitempty"`

	// 分期金额的官网价。 说明：暂只支持ECS产品。
	InstallmentOfficialWebsiteAmount *string `json:"installment_official_website_amount,omitempty"`

	// 分期付款的周期类型。 2：月。说明：暂只支持ECS产品。
	InstallmentPeriodType *int32 `json:"installment_period_type,omitempty"`
}

func (o OfficialWebsiteRatingResultV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfficialWebsiteRatingResultV2 struct{}"
	}

	return strings.Join([]string{"OfficialWebsiteRatingResultV2", string(data)}, " ")
}
