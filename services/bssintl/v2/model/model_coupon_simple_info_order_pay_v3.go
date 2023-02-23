package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CouponSimpleInfoOrderPayV3 struct {

	// 优惠券ID。
	Id string `json:"id"`

	// 优惠券类型：300:折扣券 301:促销代金券 302:促销现金券 303:促销储值卡
	Type int32 `json:"type"`
}

func (o CouponSimpleInfoOrderPayV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CouponSimpleInfoOrderPayV3 struct{}"
	}

	return strings.Join([]string{"CouponSimpleInfoOrderPayV3", string(data)}, " ")
}
