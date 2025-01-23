package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSubCustomerBillDetailRequest Request Object
type ListSubCustomerBillDetailRequest struct {

	// 语言。忽略大小写，默认 zh_cn：中文 en_us：英文
	XLanguage *string `json:"X-Language,omitempty"`

	// 账期所在月份。格式：YYYY-MM
	BillCycle string `json:"bill_cycle"`

	// 客户账号ID。您可以调用[查询客户列表](https://support.huaweicloud.com/api-bpconsole/mc_00021.html)接口获取customer_id。
	CustomerId string `json:"customer_id"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用[查询云服务类型列表](https://support.huaweicloud.com/api-bpconsole/zh-cn_topic_0000001256679455.html)接口获取。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串或携带值为null时，作为筛选条件。
	ServiceTypeCode *string `json:"service_type_code,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见[地区和终端节点](https://developer.huaweicloud.com/endpoint)对应云服务的“区域”列的值。此参数不携带或携带值为空时，不作为筛选条件；携带值为空串或携带值为null时，作为筛选条件。
	RegionCode *string `json:"region_code,omitempty"`

	// 计费模式。1：包周期3：按需10：预留实例11：节省计划此参数不携带或携带值为空或携带为null时，默认查询所有计费模式下的消费记录；不支持携带值为空串。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 账单类型。1：消费-新购 2：消费-续订 3：消费-变更 4：退款-退订 5：消费-使用 8：消费-自动续订 9：调账-补偿 14：消费-服务支持计划月末扣费 16：调账-扣费 18：消费-按月付费 20：退款-变更 23：消费-节省计划抵扣 24：退款-包年/包月转按需 103：消费-按年付费 此参数不携带或携带值为空或携带值为null时，不作为筛选条件；不支持携带值为空串。
	BillDetailType *int32 `json:"bill_detail_type,omitempty"`

	// 资源标识。此参数不携带或携带值为空时，不作为筛选条件；携带值为null时，作为筛选条件；不支持携带值为空串。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称。此参数不携带或携带值为空时，不作为筛选条件；携带值为null时，作为筛选条件；不支持携带值为空串。
	ResourceName *string `json:"resource_name,omitempty"`

	// 订单ID或交易ID，扣费维度的唯一标识。账单类型为1，2，3，4，8时为订单ID。其它场景下为交易ID。非月末扣费：应收ID月末扣费：账单ID 此参数不携带或携带值为空时，不作为筛选条件；携带值为null时，作为筛选条件；不支持携带值为空串。
	TradeId *string `json:"trade_id,omitempty"`

	// 客户经理标识。此参数不携带或携带值为空时，不作为筛选条件；携带值为null时，作为筛选条件；不支持携带值为空串。
	AccountManagerId *string `json:"account_manager_id,omitempty"`

	// 子客户的关联类型：1：顾问销售2：代售 此参数不携带或携带值为空时，不作为筛选条件；不支持携带为null和空串。
	AssociationType *string `json:"association_type,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 云经销商ID。获取方法请参见[查询云经销商列表](https://support.huaweicloud.com/api-bpconsole/espp_00003.html)。  说明： 华为云总经销商可以查询名下所有子客户消费（包括云经销商子客户）。如果是普通经销商，那么此处可以为空。如果华为云总经销商需要查询客户在云经销商关联期间的消费，需要携带该字段；除此之外，此参数不做处理。否则只能查询该客户在与自己关联期间的消费。
	IndirectPartnerId *string `json:"indirect_partner_id,omitempty"`

	// 查询的资源消费记录的开始日期，格式为YYYY-MM-DD。此参数不携带或携带值为空时，不作为筛选条件；不支持携带值为空串。 说明： 必须和bill_cycle（即资源的消费账期）在同一个月。bill_date_begin需小于等于bill_date_end。
	BillDateBegin *string `json:"bill_date_begin,omitempty"`

	// 查询的资源消费记录的结束日期，格式为YYYY-MM-DD。此参数不携带或携带值为空时，不作为筛选条件；不支持携带值为空串。 说明： 必须和bill_cycle（即资源的消费账期）在同一个月。bill_date_begin和bill_date_end两个参数必须同时出现，否则仅按照bill_cycle（即资源的消费账期）进行查询。bill_date_begin需小于等于bill_date_end。
	BillDateEnd *string `json:"bill_date_end,omitempty"`
}

func (o ListSubCustomerBillDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSubCustomerBillDetailRequest struct{}"
	}

	return strings.Join([]string{"ListSubCustomerBillDetailRequest", string(data)}, " ")
}
