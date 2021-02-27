package model

import (
	"encoding/json"

	"strings"
)

type DiscountSimpleInfo struct {
	// 折扣ID。
	Id string `json:"id"`
	// 折扣类型： 1：合同折扣2：商务优惠3：合作伙伴授予折扣609：订单调价折扣
	Type int32 `json:"type"`
}

func (o DiscountSimpleInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DiscountSimpleInfo struct{}"
	}

	return strings.Join([]string{"DiscountSimpleInfo", string(data)}, " ")
}