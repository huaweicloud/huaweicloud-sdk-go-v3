package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloudStorageAssignment 存储声明关联作用对象。
type CloudStorageAssignment struct {

	// id。
	CloudStorageAssignmentId *string `json:"cloud_storage_assignment_id,omitempty"`

	// 用户名称(单位/B)。
	AttachName *string `json:"attach_name,omitempty"`

	// 已用容量(单位/B)。
	UsedQuota *int64 `json:"used_quota,omitempty"`

	// 总容量(单位/B)。
	Quota *int64 `json:"quota,omitempty"`

	// 添加时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 云存储于本地设备的权限，上传： true : 允许上传。 false: 不允许上传。
	ActionPut *bool `json:"action_put,omitempty"`

	// 云存储于本地设备的权限，下载： true : 允许下载。 false: 不允许下载。
	ActionGet *bool `json:"action_get,omitempty"`

	// 云存储于云桌面/云应用权限，上传： true : 允许上传。 false: 不允许上传。
	RoamActionPut *bool `json:"roam_action_put,omitempty"`

	// 云存储于云桌面/云应用权限，下载： true : 允许下载。 false: 不允许下载。
	RoamActionGet *bool `json:"roam_action_get,omitempty"`
}

func (o CloudStorageAssignment) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudStorageAssignment struct{}"
	}

	return strings.Join([]string{"CloudStorageAssignment", string(data)}, " ")
}
