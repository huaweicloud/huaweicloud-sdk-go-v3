package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourcesFilters 资源计数过滤器。
type ResourcesFilters struct {

	// 帐号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`

	// 资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 资源名称。
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o ResourcesFilters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcesFilters struct{}"
	}

	return strings.Join([]string{"ResourcesFilters", string(data)}, " ")
}
