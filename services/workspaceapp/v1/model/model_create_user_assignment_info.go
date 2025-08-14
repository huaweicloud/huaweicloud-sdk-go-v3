package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateUserAssignmentInfo 创建个人文件夹。
type CreateUserAssignmentInfo struct {

	// 目标用户id。
	AttachId string `json:"attach_id"`

	// 目标用户。
	AttachName string `json:"attach_name"`

	// 配额。
	FileSystemQuota *int64 `json:"file_system_quota,omitempty"`

	// 云存储于本地设备的权限，上传： true : 允许上传。 false: 不允许上传。
	ActionPut *bool `json:"action_put,omitempty"`

	// 云存储于本地设备的权限，下载： true : 允许下载。 false: 不允许下载。
	ActionGet *bool `json:"action_get,omitempty"`

	// 云存储于云桌面/云应用权限，上传： true : 允许上传。 false: 不允许上传。
	RoamActionPut *bool `json:"roam_action_put,omitempty"`

	// 云存储于云桌面/云应用权限，下载： true : 允许下载。 false: 不允许下载。
	RoamActionGet *bool `json:"roam_action_get,omitempty"`
}

func (o CreateUserAssignmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateUserAssignmentInfo struct{}"
	}

	return strings.Join([]string{"CreateUserAssignmentInfo", string(data)}, " ")
}
