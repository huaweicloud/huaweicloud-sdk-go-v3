package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DiscountSimpleInfoV3 struct {

	// 折扣ID。
	Id string `json:"id"`

	// 折扣类型： 0：促销折扣2：商务优惠3：合作伙伴授予折扣 说明： 订单支付时，如果包含609折扣（订单调价折扣），则必须使用，不能再更换折扣类型。
	Type int32 `json:"type"`
}

func (o DiscountSimpleInfoV3) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiscountSimpleInfoV3 struct{}"
	}

	return strings.Join([]string{"DiscountSimpleInfoV3", string(data)}, " ")
}
