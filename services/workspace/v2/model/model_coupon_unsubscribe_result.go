package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CouponUnsubscribeResult 券的退订金额（降配存在）
type CouponUnsubscribeResult struct {

	// 券ID
	CouponId *string `json:"coupon_id,omitempty"`

	// 券类型 302：现金券 303：储值卡
	CouponType *string `json:"coupon_type,omitempty"`

	// 券退的金额
	Amount *float64 `json:"amount,omitempty"`

	// 度量单位'
	MeasureId *int32 `json:"measure_id,omitempty"`
}

func (o CouponUnsubscribeResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponUnsubscribeResult struct{}"
	}

	return strings.Join([]string{"CouponUnsubscribeResult", string(data)}, " ")
}
