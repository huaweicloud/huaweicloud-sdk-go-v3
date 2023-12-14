package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RoleDeployMeta 角色部署元数据
type RoleDeployMeta struct {

	// 其他扩展属性
	Other map[string]interface{} `json:"other,omitempty"`

	// 角色名称
	Name *string `json:"name,omitempty"`

	// 角色简称
	CodeName *string `json:"code_name,omitempty"`

	// 角色所属组件
	Component *string `json:"component,omitempty"`

	// 部署倾向
	NodePreference *string `json:"node_preference,omitempty"`

	// 角色数量限制
	Count *string `json:"count,omitempty"`

	// 亲和
	Affinity *string `json:"affinity,omitempty"`

	// 亲和目标
	AffinityTarget *string `json:"affinity_target,omitempty"`

	// 多实例
	MultiInstance *int32 `json:"multi_instance,omitempty"`

	// 角色类型
	RoleKind *string `json:"role_kind,omitempty"`

	// 角色限制，包含当前组件角色的一些功能限制，例如：\"no_scale_in\"
	Constraints *[]string `json:"constraints,omitempty"`

	// 多az部署
	MultiAzPlacement *string `json:"multi_az_placement,omitempty"`

	// 仲裁部署
	ArbitrationDeployment *bool `json:"arbitration_deployment,omitempty"`

	// 支持ELB
	SupportElb *bool `json:"support_elb,omitempty"`

	// 启用多亲和组
	MultiAffinityGroupEnable *bool `json:"multi_affinity_group_enable,omitempty"`

	// 本地盘反亲和
	LocalDisksAntiAffinity *bool `json:"local_disks_anti_affinity,omitempty"`

	// 多实例名称模式
	MultiInstanceNamePattern *string `json:"multi_instance_name_pattern,omitempty"`

	// 私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 权重
	Weight *string `json:"weight,omitempty"`
}

func (o RoleDeployMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleDeployMeta struct{}"
	}

	return strings.Join([]string{"RoleDeployMeta", string(data)}, " ")
}
