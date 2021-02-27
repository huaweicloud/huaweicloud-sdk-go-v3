package model

import (
	"encoding/json"

	"strings"
)

type QueryResRecordsDetailReq struct {
	// 查询的资源详单所在账期，格式为YYYY-MM。 示例：2019-01。  说明： 不支持2019年1月份之前的资源详单。
	Cycle string `json:"cycle"`
	// 云服务类型编码，例如ECS的云服务类型编码为“hws.service.type.ec2”。您可以调用查询云服务类型列表接口获取。
	CloudServiceType *string `json:"cloud_service_type,omitempty"`
	// 资源类型编码，例如ECS的VM为“hws.resource.type.vm”。您可以调用查询资源类型列表接口获取。
	ResourceType *string `json:"resource_type,omitempty"`
	// 云服务区编码，例如：“cn-north-1”。具体请参见地区和终端节点对应云服务的“区域”列的值。
	Region *string `json:"region,omitempty"`
	// 资源实例ID。
	ResInstanceId *string `json:"res_instance_id,omitempty"`
	// 计费模式： 1 : 包年/包月3：按需10：预留实例
	ChargeMode *int32 `json:"charge_mode,omitempty"`
	// 账单类型： 1：消费-新购2：消费-续订3：消费-变更4：退款-退订5：消费-使用8：消费-自动续订9：调账-补偿14：消费-服务支持计划月末扣费16：调账-扣费
	BillType *int32 `json:"bill_type,omitempty"`
	// 企业项目标识（企业项目ID）。 default项目对应ID：0未归集（表示该云服务不支持企业项目管理能力）项目对应ID：-1其余项目对应ID获取方法请参见如何获取企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 返回是否包含应付金额为0的记录。 true: 包含false: 不包含
	IncludeZeroRecord *bool `json:"include_zero_record,omitempty"`
	// 偏移量，从0开始。默认值为0。
	Offset *int32 `json:"offset,omitempty"`
	// 页面大小。默认值为10。
	Limit *int32 `json:"limit,omitempty"`
	// 查询方式。 oneself：客户自己sub_customer：企业子客户all：客户自己和企业子客户 默认为all，如果没有企业子客户，all的时候也是查询客户自己的数据。
	Method *string `json:"method,omitempty"`
	// 企业子账号ID。  说明： 如果method取值不为sub_customer，则该参数无效。如果method取值为sub_customer，则该参数不能为空。
	SubCustomerId *string `json:"sub_customer_id,omitempty"`
}

func (o QueryResRecordsDetailReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QueryResRecordsDetailReq struct{}"
	}

	return strings.Join([]string{"QueryResRecordsDetailReq", string(data)}, " ")
}
