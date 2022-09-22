package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponSimpleInfoOrderPay struct {

	// 优惠券ID。
	Id string `json:"id"`

	// 折扣类型：300:折扣券 301:促销代金券 302:促销现金券
	Type int32 `json:"type"`
}

func (o CouponSimpleInfoOrderPay) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponSimpleInfoOrderPay struct{}"
	}

	return strings.Join([]string{"CouponSimpleInfoOrderPay", string(data)}, " ")
}
