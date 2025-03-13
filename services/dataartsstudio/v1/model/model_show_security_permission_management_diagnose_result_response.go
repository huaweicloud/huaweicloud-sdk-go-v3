package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityPermissionManagementDiagnoseResultResponse Response Object
type ShowSecurityPermissionManagementDiagnoseResultResponse struct {

	// 诊断任务id。
	TaskId *string `json:"task_id,omitempty"`

	// 最新检测时间。
	CheckTime *int64 `json:"check_time,omitempty"`

	// 是否正在诊断。
	Scanning *bool `json:"scanning,omitempty"`

	HighPermission *HighPermission `json:"high_permission,omitempty"`

	UnreasonablePermission *UnreasonablePermission `json:"unreasonable_permission,omitempty"`
	HttpStatusCode         int                     `json:"-"`
}

func (o ShowSecurityPermissionManagementDiagnoseResultResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPermissionManagementDiagnoseResultResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityPermissionManagementDiagnoseResultResponse", string(data)}, " ")
}
