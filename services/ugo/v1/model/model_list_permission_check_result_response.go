package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListPermissionCheckResultResponse struct {

	// 权限检查的总条目个数。
	TotalCount *int32 `json:"total_count,omitempty" xml:"total_count"`

	// 权限检查的通过条目。
	PassedPermissionItems *[]PermissionItem `json:"passed_permission_items,omitempty" xml:"passed_permission_items"`

	// 权限检查的告警条目。
	AlarmPermissionItems *[]PermissionItem `json:"alarm_permission_items,omitempty" xml:"alarm_permission_items"`

	// 权限检查的失败条目。
	FailedPermissionItems *[]PermissionItem `json:"failed_permission_items,omitempty" xml:"failed_permission_items"`

	// 权限检查的通过条目个数。
	PassedCount *int32 `json:"passed_count,omitempty" xml:"passed_count"`

	// 权限检查的告警条目个数。
	AlarmCount *int32 `json:"alarm_count,omitempty" xml:"alarm_count"`

	// 权限检查的失败条目个数。
	FailedCount    *int32 `json:"failed_count,omitempty" xml:"failed_count"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPermissionCheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionCheckResultResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionCheckResultResponse", string(data)}, " ")
}
