package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComponentInfo 组件实例信息。
type ComponentInfo struct {

	// 组件ID。
	Id *string `json:"id,omitempty"`

	// 组件名。
	Name *string `json:"name,omitempty"`

	// 组件所在组名称。
	InstanceGroupName *string `json:"instance_group_name,omitempty"`

	// 运行状态。
	RunningStatus *string `json:"running_status,omitempty"`

	// HA状态。
	HaStatus *string `json:"ha_status,omitempty"`

	// 配置状态。
	ConfigStatus *string `json:"config_status,omitempty"`

	// 角色。
	RoleName *string `json:"role_name,omitempty"`

	// 角色缩写。
	RoleShortName *string `json:"role_short_name,omitempty"`

	// 角色类型。
	RoleType *string `json:"role_type,omitempty"`

	// 服务名。
	ServiceName *string `json:"service_name,omitempty"`

	// 对名。
	PairName *string `json:"pair_name,omitempty"`

	// 关联对。
	RelationPairs *string `json:"relation_pairs,omitempty"`

	// 是否支持Decom。
	SupportDecom *string `json:"support_decom,omitempty"`

	// 是否支持重装。
	SupportReinstall *string `json:"support_reinstall,omitempty"`

	// 是否支持收集堆栈信息。
	SupportCollectStackInfo *string `json:"support_collect_stack_info,omitempty"`
}

func (o ComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentInfo struct{}"
	}

	return strings.Join([]string{"ComponentInfo", string(data)}, " ")
}
