package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 资源
type PolicyResource struct {

	// 资源id
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 云服务名称
	ResourceProvider *string `json:"resource_provider,omitempty"`

	// 资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 区域id
	RegionId *string `json:"region_id,omitempty"`

	// 资源所属用户ID
	DomainId *string `json:"domain_id,omitempty"`
}

func (o PolicyResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyResource struct{}"
	}

	return strings.Join([]string{"PolicyResource", string(data)}, " ")
}
