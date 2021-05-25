package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListCustomerOrdersRequest struct {
	// 订单ID。

	OrderId *string `json:"order_id,omitempty"`
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId *string `json:"customer_id,omitempty"`
	// 订单创建开始时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。

	CreateTimeBegin *string `json:"create_time_begin,omitempty"`
	// 订单创建结束时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。

	CreateTimeEnd *string `json:"create_time_end,omitempty"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// 订单状态。 1：待审核3：处理中4：已取消5：已完成6：待支付9：待确认

	Status *int32 `json:"status,omitempty"`
	// 订单类型。 1：开通2：续订3：变更4：退订11：按需转包年/包月13：试用14：转商用15：费用调整

	OrderType *string `json:"order_type,omitempty"`
	// 每次查询的订单数量，默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 偏移量，从0开始。默认值为0。

	Offset *int32 `json:"offset,omitempty"`
	// 查询的订单列表排序。 支持按照创建时间进行排序，带-表示倒序。 创建时间：升序为createTime，倒序为-createTime。

	OrderBy *string `json:"order_by,omitempty"`
	// 订单支付开始时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。

	PaymentTimeBegin *string `json:"payment_time_begin,omitempty"`
	// 订单支付结束时间。 UTC时间，格式：yyyy-MM-dd'T'HH:mm:ss'Z'，如“2019-05-06T08:05:01Z”。 其中，HH范围是0～23，mm和ss范围是0～59。

	PaymentTimeEnd *string `json:"payment_time_end,omitempty"`
	// 精英服务商ID。 华为云伙伴能力中心（一级经销商）查询精英服务商的客户订单列表时，需要携带该参数；否则只能查询自己客户的订单列表。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
}

func (o ListCustomerOrdersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCustomerOrdersRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerOrdersRequest", string(data)}, " ")
}
