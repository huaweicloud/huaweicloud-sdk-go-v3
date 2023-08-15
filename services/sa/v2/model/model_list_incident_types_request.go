package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListIncidentTypesRequest Request Object
type ListIncidentTypesRequest struct {

	// 工作空间ID
	WorkspaceId string `json:"workspace_id"`

	// 子类类型标识码
	ParentBusinessCode *string `json:"parent_business_code,omitempty"`

	// request offset, from 0
	Offset *int32 `json:"offset,omitempty"`

	// request limit size
	Limit *int32 `json:"limit,omitempty"`

	// sort order, ASC, DESC.
	Order *string `json:"order,omitempty"`

	// sort by property, create_time.
	Sortby *string `json:"sortby,omitempty"`

	// name
	Name *string `json:"name,omitempty"`

	// 是否启用
	Enabled *bool `json:"enabled,omitempty"`

	// 布局名称
	LayoutName *string `json:"layout_name,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`
}

func (o ListIncidentTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListIncidentTypesRequest struct{}"
	}

	return strings.Join([]string{"ListIncidentTypesRequest", string(data)}, " ")
}
