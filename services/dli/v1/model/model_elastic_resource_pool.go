package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ElasticResourcePool 弹性资源池列表
type ElasticResourcePool struct {

	// 资源池名称
	ElasticResourcePoolName *string `json:"elastic_resource_pool_name,omitempty"`

	// 资源池id
	Id *int32 `json:"id,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	Queues *[]string `json:"queues,omitempty"`

	// 用户名
	Owner *string `json:"owner,omitempty"`

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

	// 弹性资源池属性字段。默认为标准版弹性资源池；{\"spec\":\"basic\"}标识基础版弹性资源池；{\"billing_spec_code\":\"developer\"}标识开发者弹性资源池。目前不支持其它属性设置。
	Label map[string]string `json:"label,omitempty"`
}

func (o ElasticResourcePool) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ElasticResourcePool struct{}"
	}

	return strings.Join([]string{"ElasticResourcePool", string(data)}, " ")
}
