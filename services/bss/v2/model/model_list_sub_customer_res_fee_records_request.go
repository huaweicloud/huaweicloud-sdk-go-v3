package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListSubCustomerResFeeRecordsRequest struct {
	// 客户账号ID。您可以调用查询客户列表接口获取customer_id。

	CustomerId string `json:"customer_id"`
	// 查询的客户消费记录所在账期，格式：YYYY-MM。

	Cycle string `json:"cycle"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	CloudServiceType *string `json:"cloud_service_type,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	Region *string `json:"region,omitempty"`
	// 计费模式。1 : 包年/包月3：按需10: 预留实例

	ChargeMode string `json:"charge_mode"`
	// 账单类型。1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费

	BillType *int32 `json:"bill_type,omitempty"`
	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数量限制。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 资源ID。

	ResourceId *string `json:"resource_id,omitempty"`
	// 返回是否包含应付金额为0的记录。true：包含false：不包含

	IncludeZeroRecord *bool `json:"include_zero_record,omitempty"`
	// 精英服务商ID。获取方法请参见查询精英服务商列表。如果华为云伙伴能力中心需要查询客户在精英服务商关联期间的消费，需要携带该字段；否则只能查询该客户在与自己关联期间的消费。

	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`
	// 查询的资源消费记录的开始日期，格式为YYYY-MM-DD。 说明： 必须和cycle（即资源的消费账期）在同一个月。

	BillDateBegin *string `json:"bill_date_begin,omitempty"`
	// 查询的资源消费记录的结束日期，格式为YYYY-MM-DD。 说明： 必须和cycle（即资源的消费账期）在同一个月。bill_date_begin和bill_date_end两个参数必须同时出现，否则仅按照cycle（即资源的消费账期）进行查询。

	BillDateEnd *string `json:"bill_date_end,omitempty"`
}

func (o ListSubCustomerResFeeRecordsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListSubCustomerResFeeRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerResFeeRecordsRequest", string(data)}, " ")
}
