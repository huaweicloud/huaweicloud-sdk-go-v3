package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PermissionItem 权限检查条目。
type PermissionItem struct {

	// 权限类型。
	PermissionType *string `json:"permission_type,omitempty"`

	// schema名称。
	SchemaName *string `json:"schema_name,omitempty"`

	// 权限描述。
	Description *string `json:"description,omitempty"`

	// 是否通过。
	Status *string `json:"status,omitempty"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty"`

	// 失败详情。
	FailedDetail *string `json:"failed_detail,omitempty"`

	// 解决建议。
	SuggestSolution *[]string `json:"suggest_solution,omitempty"`
}

func (o PermissionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionItem struct{}"
	}

	return strings.Join([]string{"PermissionItem", string(data)}, " ")
}
