package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 权限检查条目。
type PermissionItem struct {

	// 权限类型。
	PermissionType *string `json:"permission_type,omitempty" xml:"permission_type"`

	// schema名称。
	SchemaName *string `json:"schema_name,omitempty" xml:"schema_name"`

	// 权限描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 是否通过。
	Status *string `json:"status,omitempty" xml:"status"`

	// 失败原因。
	FailedReason *string `json:"failed_reason,omitempty" xml:"failed_reason"`

	// 失败详情。
	FailedDetail *string `json:"failed_detail,omitempty" xml:"failed_detail"`

	// 解决建议。
	SuggestSolution *[]string `json:"suggest_solution,omitempty" xml:"suggest_solution"`
}

func (o PermissionItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PermissionItem struct{}"
	}

	return strings.Join([]string{"PermissionItem", string(data)}, " ")
}
