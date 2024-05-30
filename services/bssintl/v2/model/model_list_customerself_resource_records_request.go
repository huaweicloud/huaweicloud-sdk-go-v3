package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCustomerselfResourceRecordsRequest Request Object
type ListCustomerselfResourceRecordsRequest struct {

	// 语言：中文：zh_CN 英文：en_US。缺省为zh_CN
	XLanguage *string `json:"X-Language,omitempty"`

	// 查询的资源消费记录所在账期，格式：YYYY-MM。
	Cycle string `json:"cycle"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。此参数不携带时，不作为筛选条件；携带值为空或携带值为空串时，作为筛选条件。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 云服务区编码，例如：“ap-southeast-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。此参数不携带时，不作为筛选条件；携带值为空或携带值为空串时，作为筛选条件。
	Region *string `json:"region,omitempty"`

	// 计费模式。1：包年/包月3：按需10：预留实例11：节省计划。此参数不携带时，不作为筛选条件。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 账单类型。1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿 12：消费-按时计费 14：消费-服务支持计划月末扣费15：消费-税金16：调账-扣费17：消费-保底差额 说明： 保底差额=客户签约保底合同后，如果没有达到保底消费，客户需要补交的费用，仅限于直销或者伙伴顾问销售类子客户，且为后付费用户。20：退款-变更100：退款-退订税金23：消费-节省计划抵扣 24：退款-包年/包月转按需101：调账-补偿税金102：调账-扣费税金。此参数不携带或携带值为空时，不作为筛选条件。
	BillType *int32 `json:"bill_type,omitempty"`

	// 偏移量，从0开始。默认值为0。 说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 每次查询的数量限制。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 资源ID。此参数不携带时，不作为筛选条件；携带值为空或携带值为空串时，作为筛选条件。
	ResourceId *string `json:"resource_id,omitempty"`

	// 企业项目标识（企业项目ID）。default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：null。此参数不携带时，不作为筛选条件；携带值为空或携带值为空串时，作为筛选条件。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 返回是否包含应付金额为0的记录。true：包含false：不包含。此参数不携带或携带值为空时，不作为筛选条件。
	IncludeZeroRecord *bool `json:"include_zero_record,omitempty"`

	// 查询资源消费记录的方式。oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户。此参数不携带或携带值为空时，默认值为“all”，如果没有企业子客户，取值为all时查询的是客户自己的资源消费记录。
	Method *string `json:"method,omitempty"`

	// 企业子账号ID。此参数携带值为空串时，不作为筛选条件。 说明： 如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。
	SubCustomerId *string `json:"sub_customer_id,omitempty"`

	// 订单ID或交易ID。账单类型为1、2、3、4和8时此处为订单ID。账单类型为其它场景时此处为交易ID，为扣费维度的唯一标识。例如非月末扣费时为应收ID；月末扣费时为账单ID。此参数不携带时，不作为筛选条件；携带值为空或携带值为空串时，作为筛选条件。
	TradeId *string `json:"trade_id,omitempty"`

	// 查询的资源消费记录的开始日期，格式为YYYY-MM-DD。 说明： 必须和cycle（即资源的消费账期）在同一个月。bill_date_begin和bill_date_end两个参数必须同时出现，否则仅按照cycle（即资源的消费账期）进行查询。
	BillDateBegin *string `json:"bill_date_begin,omitempty"`

	// 查询的资源消费记录的结束日期，格式为YYYY-MM-DD。 说明： 必须和cycle（即资源的消费账期）在同一个月。bill_date_begin和bill_date_end两个参数必须同时出现，否则仅按照cycle（即资源的消费账期）进行查询。
	BillDateEnd *string `json:"bill_date_end,omitempty"`

	// |参数名称：统计类型。| |参数的约束及描述：统计类型。非必填，默认值为3。1：按账期 3：按明细|
	StatisticType *int32 `json:"statistic_type,omitempty"`
}

func (o ListCustomerselfResourceRecordsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCustomerselfResourceRecordsRequest struct{}"
	}

	return strings.Join([]string{"ListCustomerselfResourceRecordsRequest", string(data)}, " ")
}
