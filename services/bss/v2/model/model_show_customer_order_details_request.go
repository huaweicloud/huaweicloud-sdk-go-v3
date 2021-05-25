package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCustomerOrderDetailsRequest struct {
	// 订单ID。 查询订单列表时系统会返回订单ID。

	OrderId string `json:"order_id"`
	// 每页大小。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 精英服务商ID。 华为云伙伴能力中心（一级经销商）查询精英服务商的客户订单详情时，需要携带该参数；否则只能查询自己客户的订单详情。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ShowCustomerOrderDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCustomerOrderDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerOrderDetailsRequest", string(data)}, " ")
}
