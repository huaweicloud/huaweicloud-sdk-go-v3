package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源对象
type ResourceEntity struct {

	// 资源id
	Id *string `json:"id,omitempty" xml:"id"`

	// 资源名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 云服务名称
	Provider *string `json:"provider,omitempty" xml:"provider"`

	// 资源类型
	Type *string `json:"type,omitempty" xml:"type"`

	// 区域id
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// Openstack中的项目id
	ProjectId *string `json:"project_id,omitempty" xml:"project_id"`

	// Openstack中的项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name"`

	// 企业项目id
	EpId *string `json:"ep_id,omitempty" xml:"ep_id"`

	// 企业项目名称
	EpName *string `json:"ep_name,omitempty" xml:"ep_name"`

	// 资源详情校验码
	Checksum *string `json:"checksum,omitempty" xml:"checksum"`

	// 资源创建时间
	Created *string `json:"created,omitempty" xml:"created"`

	// 资源更新时间
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 资源操作状态
	ProvisioningState *string `json:"provisioning_state,omitempty" xml:"provisioning_state"`

	// 资源Tag
	Tags map[string]string `json:"tags,omitempty" xml:"tags"`

	// 资源详细属性
	Properties map[string]interface{} `json:"properties,omitempty" xml:"properties"`
}

func (o ResourceEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceEntity struct{}"
	}

	return strings.Join([]string{"ResourceEntity", string(data)}, " ")
}
