package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VersionComponent 组件详情
type VersionComponent struct {

	// 其他
	Other map[string]interface{} `json:"other,omitempty"`

	// 组件名称
	Name *string `json:"name,omitempty"`

	// 支持版本
	Version *string `json:"version,omitempty"`

	// 组件依赖项
	DependOn *[]string `json:"depend_on,omitempty"`

	// 组件描述
	Description *string `json:"description,omitempty"`

	// 支持该组件的集群类型
	AvailableClusterTypes *[]string `json:"available_cluster_types,omitempty"`

	// 外部数据源
	ExternalDatasources *[]ComponentExternalDatasource `json:"external_datasources,omitempty"`

	// 所需的ip等资源说明
	ResourceRequirement *[]string `json:"resource_requirement,omitempty"`

	// 有效角色
	ValidRoles *[]string `json:"valid_roles,omitempty"`

	// 是否可见
	Visible *bool `json:"visible,omitempty"`

	// 子组件
	ChildrenComponents *[]string `json:"children_components,omitempty"`

	// 多az支持状态
	MultiAzSupportStatus *string `json:"multi_az_support_status,omitempty"`
}

func (o VersionComponent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionComponent struct{}"
	}

	return strings.Join([]string{"VersionComponent", string(data)}, " ")
}
