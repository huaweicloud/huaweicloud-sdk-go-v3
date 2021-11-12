package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListCustomerselfResourceRecordsRequest struct {
	// |参数名称：语言| |参数的约束及描述：中文：zh_CN 英文：en_US。缺省为zh_CN|

	XLanguage *string `json:"X-Language,omitempty"`
	// 查询的资源消费记录所在账期，格式：YYYY-MM。

	Cycle string `json:"cycle"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。

	CloudServiceType *string `json:"cloud_service_type,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。

	Region *string `json:"region,omitempty"`
	// 计费模式。1：包年/包月3：按需10：预留实例

	ChargeMode *string `json:"charge_mode,omitempty"`
	// 账单类型。1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费

	BillType *int32 `json:"bill_type,omitempty"`
	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。

	Offset *int32 `json:"offset,omitempty"`
	// 每次查询的数量限制。默认值为10。

	Limit *int32 `json:"limit,omitempty"`
	// 资源ID。

	ResourceId *string `json:"resource_id,omitempty"`
	// 企业项目标识（企业项目ID）。default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见如何获取企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 返回是否包含应付金额为0的记录。true：包含false：不包含

	IncludeZeroRecord *bool `json:"include_zero_record,omitempty"`
	// 查询资源消费记录的方式。oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户默认为all，如果没有企业子客户，取值为all时查询的是客户自己的资源消费记录。

	Method *string `json:"method,omitempty"`
	// 企业子账号ID。 说明： 如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。

	SubCustomerId *string `json:"sub_customer_id,omitempty"`
	// 订单ID或交易ID。账单类型为1、2、3、4和8时此处为订单ID。账单类型为其它场景时此处为交易ID，为扣费维度的唯一标识。例如非月末扣费时为应收ID；月末扣费时为账单ID。

	TradeId *string `json:"trade_id,omitempty"`
	// 查询的资源消费记录的开始日期，格式为YYYY-MM-DD。 说明： 必须和cycle（即资源的消费账期）在同一个月。bill_date_begin和bill_date_end两个参数必须同时出现，否则仅按照cycle（即资源的消费账期）进行查询。

	BillDateBegin *string `json:"bill_date_begin,omitempty"`
	// 查询的资源消费记录的结束日期，格式为YYYY-MM-DD。 说明： 必须和cycle（即资源的消费账期）在同一个月。bill_date_begin和bill_date_end两个参数必须同时出现，否则仅按照cycle（即资源的消费账期）进行查询。

	BillDateEnd *string `json:"bill_date_end,omitempty"`
}

func (o ListCustomerselfResourceRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerselfResourceRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerselfResourceRecordsRequest", string(data)}, " ")
}
