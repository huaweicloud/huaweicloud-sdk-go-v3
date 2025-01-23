package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type OfficialWebsiteRatingResult struct {

	// 包年/包月产品的官网价。
	OfficialWebsiteAmount *decimal.Decimal `json:"official_website_amount,omitempty"`

	// 分期金额的官网价。  说明： 暂只支持CloudPond产品。
	InstallmentOfficialWebsiteAmount *string `json:"installment_official_website_amount,omitempty"`

	// 分期付款的周期类型。 2：月 3：年  说明： 暂只支持CloudPond产品。
	InstallmentPeriodType *int32 `json:"installment_period_type,omitempty"`

	// 价格度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 产品询价结果，具体参见表5。
	ProductRatingResults *[]PeriodProductOfficialRatingResult `json:"product_rating_results,omitempty"`
}

func (o OfficialWebsiteRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OfficialWebsiteRatingResult struct{}"
	}

	return strings.Join([]string{"OfficialWebsiteRatingResult", string(data)}, " ")
}
