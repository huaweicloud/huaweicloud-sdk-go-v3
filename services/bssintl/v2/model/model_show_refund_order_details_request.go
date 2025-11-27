package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRefundOrderDetailsRequest Request Object
type ShowRefundOrderDetailsRequest struct {

	// 退订订单或者降配订单的ID。
	OrderId string `json:"order_id"`

	// 客户账号ID，非必填，范围限制:0-64，伙伴查询子客户退款订单的金额详情必须携带该字段。
	CustomerId *string `json:"customer_id,omitempty"`

	// 云经销商ID，非必填，范围限制:0-64，如果华为云总经销商（一级经销商）需要查询云经销商的子客户退款订单的金额详情必须携带该字段。除此之外，此参数不做处理。|
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ShowRefundOrderDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRefundOrderDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowRefundOrderDetailsRequest", string(data)}, " ")
}
