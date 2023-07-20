package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ElasticResourcePoolsResponse 弹性资源池列表
type ElasticResourcePoolsResponse struct {

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	//
	Queues *[]string `json:"queues,omitempty"`

	// 用户名
	Owner *string `json:"owner,omitempty"`

	// 资源池名称
	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`

	// 资源池描述
	Description *string `json:"description,omitempty"`

	// 最大cu数量
	MaxCu *int32 `json:"max_cu,omitempty"`

	// 最小cu数量
	MinCu *int32 `json:"min_cu,omitempty"`

	// 实际cu数量
	ActualCu *int32 `json:"actual_cu,omitempty"`

	// 子网
	CidrInVpc *string `json:"cidr_in_vpc,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 当前cu数量
	CurrentCu *int32 `json:"current_cu,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// 资源ID
	ResourceId *string `json:"resource_id,omitempty"`

	// 创建失败原因
	FailReason *string `json:"fail_reason,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 预付费cu数量
	PrepayCu *int32 `json:"prepay_cu,omitempty"`

	// 计费类型
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 弹性资源池类型
	Manager *string `json:"manager,omitempty"`

	// 标识弹性资源池，目前只支持 开发者标签，\"label\": {\"billing_spec_code\":\"developer\"}
	Label map[string]string `json:"label,omitempty"`
}

func (o ElasticResourcePoolsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElasticResourcePoolsResponse struct{}"
	}

	return strings.Join([]string{"ElasticResourcePoolsResponse", string(data)}, " ")
}
