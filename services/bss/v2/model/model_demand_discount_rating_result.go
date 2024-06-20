package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/shopspring/decimal"

	"strings"
)

type DemandDiscountRatingResult struct {

	// 优惠标识ID。
	DiscountId *string `json:"discount_id,omitempty"`

	// 折扣优惠类型。 合同商务折扣：605：华为云BE场景下的合同商务折扣606：分销商BE场景下的合同商务折扣 伙伴授予折扣：607：合作伙伴授予折扣-折扣率
	DiscountType *int32 `json:"discount_type,omitempty"`

	// 折扣的金额。
	Amount *decimal.Decimal `json:"amount,omitempty"`

	// 度量单位标识。 1：元
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 折扣名称。
	DiscountName *string `json:"discount_name,omitempty"`
}

func (o DemandDiscountRatingResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DemandDiscountRatingResult struct{}"
	}

	return strings.Join([]string{"DemandDiscountRatingResult", string(data)}, " ")
}
