package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopVersionInfo 桌面版本信息。
type DesktopVersionInfo struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面sid。
	Sid *string `json:"sid,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 用户名。
	Username *string `json:"username,omitempty"`

	// 桌面状态。
	Status *string `json:"status,omitempty"`

	// 桌面执行任务状态。
	TaskStatus *string `json:"task_status,omitempty"`
}

func (o DesktopVersionInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopVersionInfo struct{}"
	}

	return strings.Join([]string{"DesktopVersionInfo", string(data)}, " ")
}
