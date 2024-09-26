package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NamespaceBasicDto namespace基本信息
type NamespaceBasicDto struct {

	// namespace id
	Id *int32 `json:"id,omitempty"`

	// namespace name
	Name *string `json:"name,omitempty"`

	// namespace path
	Path *string `json:"path,omitempty"`

	// namespace 开发模式
	DevelopMode *string `json:"develop_mode,omitempty"`

	// namespace region
	Region *string `json:"region,omitempty"`

	// namespace cell
	Cell *string `json:"cell,omitempty"`

	// namespace kind
	Kind *string `json:"kind,omitempty"`

	// namespace full path
	FullPath *string `json:"full_path,omitempty"`

	// namespace full name
	FullName *string `json:"full_name,omitempty"`

	// namespace parent id
	ParentId *int32 `json:"parent_id,omitempty"`

	// namespace visibility level
	VisibilityLevel *int32 `json:"visibility_level,omitempty"`

	// namespace file control enable
	EnableFileControl *bool `json:"enable_file_control,omitempty"`

	// namespace owner id
	OwnerId *int32 `json:"owner_id,omitempty"`
}

func (o NamespaceBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NamespaceBasicDto struct{}"
	}

	return strings.Join([]string{"NamespaceBasicDto", string(data)}, " ")
}
