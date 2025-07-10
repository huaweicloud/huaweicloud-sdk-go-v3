package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchResInquiryResult struct {

	// ID标识,同一次询价中不能重复。
	Id *string `json:"id,omitempty"`

	// 变更后的产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 总额，即最终优惠后的金额(降配场景下包含退还的现金、现金券、储值卡的总额)。
	Amount float32 `json:"amount,omitempty"`

	// 券的退订金额（降配存在）。
	CouponResults *[]CouponUnsubscribeResult `json:"coupon_results,omitempty"`

	// 优惠额。
	DiscountAmount float32 `json:"discount_amount,omitempty"`

	// 原总额，即优惠前总额。
	OriginalAmount float32 `json:"original_amount,omitempty"`

	// 度量单位标识。
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 扩展参数。
	ExtendParams *string `json:"extend_params,omitempty"`
}

func (o BatchResInquiryResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchResInquiryResult struct{}"
	}

	return strings.Join([]string{"BatchResInquiryResult", string(data)}, " ")
}
