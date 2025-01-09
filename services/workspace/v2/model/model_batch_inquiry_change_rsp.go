package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchInquiryChangeRsp struct {

	// 返回码，恒为0
	RetCode *string `json:"ret_code,omitempty"`

	// 返回信息
	ErrorTxt *string `json:"error_txt,omitempty"`

	// 总额，即最终优惠后的金额(降配场景下包含退还的现金、现金券、储值卡的总额)
	Amount float32 `json:"amount,omitempty"`

	// 券的退订金额（降配存在）
	CouponResults *[]CouponUnsubscribeResult `json:"coupon_results,omitempty"`

	// 优惠额
	DiscountAmount float32 `json:"discount_amount,omitempty"`

	// 原总额，即优惠前总额
	OriginalAmount float32 `json:"original_amount,omitempty"`

	// 度量单位标识
	MeasureId *int32 `json:"measure_id,omitempty"`

	// 货币单位代码（遵循ISO 4217标准）
	Currency *string `json:"currency,omitempty"`

	// 批价结果
	ProductRatingResult *[]BatchResInquiryResult `json:"product_rating_result,omitempty"`

	// 扩展参数
	ExtendParams *string `json:"extend_params,omitempty"`
}

func (o BatchInquiryChangeRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchInquiryChangeRsp struct{}"
	}

	return strings.Join([]string{"BatchInquiryChangeRsp", string(data)}, " ")
}
