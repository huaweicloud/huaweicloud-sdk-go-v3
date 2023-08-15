package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DataClassTypeDetailInfo Info of Dataclass Type
type DataClassTypeDetailInfo struct {

	// Id value
	Id *string `json:"id,omitempty"`

	// The name, display only
	ParentName *string `json:"parent_name,omitempty"`

	// 类型标识码
	ParentBusinessCode *string `json:"parent_business_code,omitempty"`

	// The name, display only
	Name *string `json:"name,omitempty"`

	// 类型标识码
	BusinessCode *string `json:"business_code,omitempty"`

	// The description, display only
	Description *string `json:"description,omitempty"`

	// workspace id
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// Project id value
	ProjectId *string `json:"project_id,omitempty"`

	// If is enabled, false for disenabled, true for enabled
	Enabled *bool `json:"enabled,omitempty"`

	// 是否内置
	IsBuiltIn *bool `json:"is_built_in,omitempty"`

	// 布局ID
	LayoutId *string `json:"layout_id,omitempty"`

	// The name, display only
	LayoutName *string `json:"layout_name,omitempty"`

	// dataclass id.
	DataclassId *string `json:"dataclass_id,omitempty"`

	// sla
	Sla *int32 `json:"sla,omitempty"`
}

func (o DataClassTypeDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DataClassTypeDetailInfo struct{}"
	}

	return strings.Join([]string{"DataClassTypeDetailInfo", string(data)}, " ")
}
