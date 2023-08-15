package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryResRecordsDetailReq struct {

	// 查询的资源详单所在账期，东八区时间，格式为YYYY-MM。 示例：2019-01。  说明： 不支持2019年1月份之前的资源详单。
	Cycle string `json:"cycle"`

	// 云服务类型编码，例如OBS的云服务类型编码为“hws.service.type.obs”。您可以调用查询云服务类型列表接口获取。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`

	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ResourceType *string `json:"resource_type,omitempty"`

	// 云服务区编码，例如：“ap-southeast-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	Region *string `json:"region,omitempty"`

	// 资源实例ID。 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	ResInstanceId *string `json:"res_instance_id,omitempty"`

	// 计费模式。 1 : 包年/包月3：按需10：预留实例 此参数不携带或者携带值为null时，返回所有计费模式的资源详单数据记录。
	ChargeMode *int32 `json:"charge_mode,omitempty"`

	// 账单类型： 1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费15：消费-税金16：调账-扣费17：消费-保底差额 说明： 保底差额=客户签约保底合同后，如果没有达到保底消费，客户需要补交的费用，仅限于直销或者伙伴顾问销售类子客户，且为后付费用户。 20：退款-变更100：退款-退订税金101：调账-补偿税金102：调账-扣费税金 此参数不携带或者携带值为null时，返回所有账单类型的资源详单数据记录。
	BillType *int32 `json:"bill_type,omitempty"`

	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：null 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 返回是否包含应付金额为0的记录。 true: 包含false: 不包含 此参数不携带或携带值为空串或携带值为null时，不作为筛选条件。
	IncludeZeroRecord *bool `json:"include_zero_record,omitempty"`

	// 偏移量，从0开始。默认值为0。  说明： offset用于分页处理，如不涉及分页，请使用默认值0。offset表示相对于满足条件的第一个数据的偏移量。如offset = 1，则返回满足条件的第二个数据至最后一个数据。 例如，满足查询条件的结果共10条数据，limit取值为10，offset取值为1，则返回的数据为2~10，第一条数据不返回。
	Offset *int32 `json:"offset,omitempty"`

	// 页面大小。默认值为10。
	Limit *int32 `json:"limit,omitempty"`

	// 查询方式。 oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户 此参数不携带或携带值为空串或携带值为null时，默认值为“all”，如果没有企业子客户，all的时候也是查询客户自己的数据。
	Method *string `json:"method,omitempty"`

	// 企业子账号ID。  说明： 如果method取值不为sub_customer，则此参数无效。如果method取值为sub_customer，则此参数不能为空。
	SubCustomerId *string `json:"sub_customer_id,omitempty"`

	// 统计类型。默认值为1。 1：按账期2：按天
	StatisticType *int32 `json:"statistic_type,omitempty"`
}

func (o QueryResRecordsDetailReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryResRecordsDetailReq struct{}"
	}

	return strings.Join([]string{"QueryResRecordsDetailReq", string(data)}, " ")
}
