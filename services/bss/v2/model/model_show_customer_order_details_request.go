package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCustomerOrderDetailsRequest struct {
	// |参数名称：语言| |参数的约束及描述：中文：zh_CN 英文：en_US缺省为zh_CN|

	XLanguage *string `json:"X-Language,omitempty"`
	// |参数名称：订单ID。| |参数的约束及描述：|

	OrderId string `json:"order_id"`
	// 每页大小。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 精英服务商ID。华为云伙伴能力中心（一级经销商）查询精英服务商的客户订单详情时，需要携带该参数；否则只能查询自己客户的订单详情。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ShowCustomerOrderDetailsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCustomerOrderDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerOrderDetailsRequest", string(data)}, " ")
}
