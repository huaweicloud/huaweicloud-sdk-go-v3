package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type OfficialWebsiteRatingResult struct {

	// 包年/包月产品的官网价。
	OfficialWebsiteAmount *float64 `json:"official_website_amount,omitempty" xml:"official_website_amount"`

	// 分期金额的官网价。  说明： 暂只支持IES产品。
	InstallmentOfficialWebsiteAmount *string `json:"installment_official_website_amount,omitempty" xml:"installment_official_website_amount"`

	// 分期付款的周期类型。 2：月  说明： 暂只支持IES产品。
	InstallmentPeriodType *int32 `json:"installment_period_type,omitempty" xml:"installment_period_type"`

	// 价格度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty" xml:"measure_id"`

	// 产品询价结果，具体参见表5。
	ProductRatingResults *[]PeriodProductOfficialRatingResult `json:"product_rating_results,omitempty" xml:"product_rating_results"`
}

func (o OfficialWebsiteRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfficialWebsiteRatingResult struct{}"
	}

	return strings.Join([]string{"OfficialWebsiteRatingResult", string(data)}, " ")
}
