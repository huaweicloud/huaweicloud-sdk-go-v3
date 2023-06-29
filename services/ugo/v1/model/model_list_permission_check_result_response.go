package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPermissionCheckResultResponse Response Object
type ListPermissionCheckResultResponse struct {

	// 权限检查的总条目个数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 权限检查的通过条目。
	PassedPermissionItems *[]PermissionItem `json:"passed_permission_items,omitempty"`

	// 权限检查的告警条目。
	AlarmPermissionItems *[]PermissionItem `json:"alarm_permission_items,omitempty"`

	// 权限检查的失败条目。
	FailedPermissionItems *[]PermissionItem `json:"failed_permission_items,omitempty"`

	// 权限检查的通过条目个数。
	PassedCount *int32 `json:"passed_count,omitempty"`

	// 权限检查的告警条目个数。
	AlarmCount *int32 `json:"alarm_count,omitempty"`

	// 权限检查的失败条目个数。
	FailedCount    *int32 `json:"failed_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListPermissionCheckResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPermissionCheckResultResponse struct{}"
	}

	return strings.Join([]string{"ListPermissionCheckResultResponse", string(data)}, " ")
}
