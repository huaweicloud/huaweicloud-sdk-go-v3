package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResRecordsDetailReq struct {

	// 查询的资源详单所在账期，东八区时间，格式为YYYY-MM。 示例：2019-01。  说明： 不支持2019年1月份之前的资源详单。
	Cycle string `json:"cycle"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	Region *string `json:"region,omitempty"`

	// 资源实例ID。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ResInstanceId *string `json:"res_instance_id,omitempty"`

	// 计费模式。 1 : 包年/包月3：按需10：预留实例11：节省计划 此参数不携带或者携带值为null时，返回所有计费模式的资源详单数据记录。
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 账单类型。1：消费-新购 2：消费-续订 3：消费-变更 4：退款-退订 5：消费-使用 8：消费-自动续订 9：调账-补偿 14：消费-服务支持计划月末扣费 16：调账-扣费 18：消费-按月付费 20：退款-变更 23：消费-节省计划抵扣 24：退款-包年/包月转按需 103：消费-按年付费 此参数不携带或者携带值为null时，返回所有账单类型的资源详单数据记录。
	BillType *int32 `json:"bill_type,omitempty"`

	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：null其余项目对应ID获取方法请参见[如何获取企业项目ID](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0126101490.html)。此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 返回是否包含应付金额为0的记录。 true: 包含false: 不包含 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	IncludeZeroRecord *bool `json:"include_zero_record,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 查询方式。 oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户 此参数不携带或携带值为空串或携带值为null时，默认值为“all”，如果没有企业子客户，all的时候也是查询客户自己的数据。  说明： 若需要查询财务独立企业子的账单信息，前提是子账号已经接受了企业主账号的“允许查看子账号消费信息”的申请。申请方法见“变更子账号权限”。
	Method *string `json:"method,omitempty"`

	// 企业子账号ID。  说明： 如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。
	SubCustomerId *string `json:"sub_customer_id,omitempty"`

	// 统计类型。默认值为1。 1：按账期2：按天3：按明细
	StatisticType *int32 `json:"statistic_type,omitempty"`

	// 查询类型。默认值为BILLCYCLE。 BILLCYCLE：按月DAILY：按天仅当statistic_type=2或3时，支持传递query_type=DAILY。该参数不携带或携带值为null或携带为空串时，取默认值BILLCYCLE。
	QueryType *string `json:"query_type,omitempty"`

	// 账期开始时间。格式为YYYY-MM-DD。 仅当query_type=DAILY时，必须传递账期开始时间。该参数不携带或携带值为null或携带为空串时，不作为筛选条件。
	BillCycleBegin *string `json:"bill_cycle_begin,omitempty"`

	// 账期结束时间。格式为YYYY-MM-DD。 仅当query_type=DAILY时，必须传递账期结束时间。该参数不携带或携带值为null或携带为空串时，不作为筛选条件。
	BillCycleEnd *string `json:"bill_cycle_end,omitempty"`

	// |参数名称：支付账号ID。| |参数的约束及描述：普通客户、财务独立企业子客户查询消费记录，只能查询到客户自己的消费记录，该参数不携带或携带为自身ID时，查询的都只是自身的消费记录； 企业主客户查询消费记录，不携带时，查询的是自身的、财务托管企业子、财务独立企业子的消费记录；入参自身ID时，查询的是自身的、财务托管企业子的消费记录；也可入参其名下财务独立企业子的客户ID，只查询该财务独立企业子的消费记录； 财务托管企业子查询消费记录，入参自身ID时，查询的是未与企业主关联时的消费记录；入参企业主客户ID时，查询的是与企业主关联后的消费记录；不携带时查询以上全部消费记录|
	PayerAccountId *string `json:"payer_account_id,omitempty"`
}

func (o QueryResRecordsDetailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResRecordsDetailReq struct{}"
	}

	return strings.Join([]string{"QueryResRecordsDetailReq", string(data)}, " ")
}
