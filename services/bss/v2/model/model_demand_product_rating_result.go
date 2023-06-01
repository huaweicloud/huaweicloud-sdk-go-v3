package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type DemandProductRatingResult struct {

	// 同一次询价中不能重复，用于标识返回询价结果和请求的映射关系。
	Id *string `json:"id,omitempty"`

	// 按需产品的ID。
	ProductId *string `json:"product_id,omitempty"`

	// 折扣的金额。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 优惠额（官网价和总价的差）。
	DiscountAmount *decimal.Decimal `json:"discount_amount,omitempty"`

	// 按需产品的官网价。
	OfficialWebsiteAmount *decimal.Decimal `json:"official_website_amount,omitempty"`

	// 度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 折扣优惠明细，包含产品本身的促销信息，同时包含商务或者伙伴折扣的优惠信息，具体参见表3。
	DiscountRatingResults *[]DemandDiscountRatingResult `json:"discount_rating_results,omitempty"`
}

func (o DemandProductRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandProductRatingResult struct{}"
	}

	return strings.Join([]string{"DemandProductRatingResult", string(data)}, " ")
}
