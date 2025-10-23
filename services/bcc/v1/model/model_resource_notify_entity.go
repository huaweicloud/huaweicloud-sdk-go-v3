package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceNotifyEntity 资源变化通知格式
type ResourceNotifyEntity struct {

	// 资源名称
	Name *string `json:"name,omitempty"`

	// 区域ID
	RegionId *string `json:"region_id,omitempty"`

	// Openstack中的项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 企业项目ID
	EpId *string `json:"ep_id,omitempty"`

	// 资源创建时间
	Created *string `json:"created,omitempty"`

	// 资源更新时间
	Updated *string `json:"updated,omitempty"`

	// 资源操作状态
	ProvisioningState *string `json:"provisioning_state,omitempty"`

	// 资源标签
	Tags map[string]string `json:"tags,omitempty"`

	// 资源详细属性
	Properties map[string]interface{} `json:"properties,omitempty"`
}

func (o ResourceNotifyEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceNotifyEntity struct{}"
	}

	return strings.Join([]string{"ResourceNotifyEntity", string(data)}, " ")
}
