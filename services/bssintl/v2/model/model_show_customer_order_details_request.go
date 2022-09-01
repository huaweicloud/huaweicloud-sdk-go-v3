package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowCustomerOrderDetailsRequest struct {

	// 语言：中文：zh_CN 英文：en_US 缺省为zh_CN|
	XLanguage *string `json:"X-Language,omitempty" xml:"X-Language"`

	// 订单ID。
	OrderId string `json:"order_id" xml:"order_id"`

	// 每页大小。默认值为10。
	Limit *int32 `json:"limit,omitempty" xml:"limit"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty" xml:"offset"`

	// 云经销商ID。华为云总经销商（一级经销商）查询云经销商的客户订单详情时，需要携带该参数；除此之外，此参数不做处理。否则只能查询自己客户的订单详情。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty" xml:"indirect_partner_id"`
}

func (o ShowCustomerOrderDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCustomerOrderDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomerOrderDetailsRequest", string(data)}, " ")
}
