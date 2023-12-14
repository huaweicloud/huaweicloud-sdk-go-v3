package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMrsVersionMetadataResponse Response Object
type ShowMrsVersionMetadataResponse struct {

	// 其他
	Other map[string]interface{} `json:"other,omitempty"`

	// 镜像版本名称
	Name *string `json:"name,omitempty"`

	// 模板类型
	TemplateType *string `json:"template_type,omitempty"`

	// 镜像ID
	ImageId *string `json:"image_id,omitempty"`

	// 版本状态
	Status *string `json:"status,omitempty"`

	// 特性列表
	Features *[]string `json:"features,omitempty"`

	// 集群类型列表
	ClusterTypes *[]string `json:"cluster_types,omitempty"`

	// 版本类型
	VersionType *string `json:"version_type,omitempty"`

	// 组件列表
	Components *[]VersionComponent `json:"components,omitempty"`

	// 版本所需的ip等资源说明
	ResourceRequirement *[]string `json:"resource_requirement,omitempty"`

	Constraints *VersionConstraint `json:"constraints,omitempty"`

	Flavors *FlavorLists `json:"flavors,omitempty"`

	// 版本组件实例角色部署策略
	RoleDeployMeta *[]RoleDeployMeta `json:"role_deploy_meta,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowMrsVersionMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMrsVersionMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowMrsVersionMetadataResponse", string(data)}, " ")
}
