package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeCloudStorageAssignmentInfo 切换集群后文件夹返回信息。
type ChangeCloudStorageAssignmentInfo struct {

	// 存储分配的唯一标识符。
	Id *string `json:"id,omitempty"`

	// 目标用户id。
	AttachId *string `json:"attach_id,omitempty"`

	// 目标用户。
	Attach *string `json:"attach,omitempty"`

	// 配额。
	Quota *int64 `json:"quota,omitempty"`

	// 云存储于本地设备的权限，上传： true : 允许上传。 false: 不允许上传。
	ActionPut *bool `json:"action_put,omitempty"`

	// 云存储于本地设备的权限，下载： true : 允许下载。 false: 不允许下载。
	ActionGet *bool `json:"action_get,omitempty"`

	// 云存储于云桌面/云应用权限，上传： true : 允许上传。 false: 不允许上传。
	RoamActionPut *bool `json:"roam_action_put,omitempty"`

	// 云存储于云桌面/云应用权限，下载： true : 允许下载。 false: 不允许下载。
	RoamActionGet *bool `json:"roam_action_get,omitempty"`
}

func (o ChangeCloudStorageAssignmentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeCloudStorageAssignmentInfo struct{}"
	}

	return strings.Join([]string{"ChangeCloudStorageAssignmentInfo", string(data)}, " ")
}
