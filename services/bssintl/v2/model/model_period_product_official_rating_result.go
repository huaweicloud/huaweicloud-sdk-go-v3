package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type PeriodProductOfficialRatingResult struct {

	// ID标识，来源于请求中的ID。
	Id *string `json:"id,omitempty"`

	// 包年/包月产品的ID。
	ProductId *string `json:"product_id,omitempty"`

	// 包年/包月产品的官网价。
	OfficialWebsiteAmount *decimal.Decimal `json:"official_website_amount,omitempty"`

	// 价格度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o PeriodProductOfficialRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodProductOfficialRatingResult struct{}"
	}

	return strings.Join([]string{"PeriodProductOfficialRatingResult", string(data)}, " ")
}
