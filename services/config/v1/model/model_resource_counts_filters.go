package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ResourceCountsFilters 资源计数过滤器。
type ResourceCountsFilters struct {

	// 帐号ID。
	AccountId *string `json:"account_id,omitempty"`

	// 资源类型。
	ResourceType *string `json:"resource_type,omitempty"`

	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`
}

func (o ResourceCountsFilters) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourceCountsFilters struct{}"
	}

	return strings.Join([]string{"ResourceCountsFilters", string(data)}, " ")
}
