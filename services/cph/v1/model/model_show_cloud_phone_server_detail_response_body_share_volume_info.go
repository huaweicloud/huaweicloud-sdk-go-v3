package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCloudPhoneServerDetailResponseBodyShareVolumeInfo 云手机服务器的共享存储相关信息。
type ShowCloudPhoneServerDetailResponseBodyShareVolumeInfo struct {

	// 共享存储磁盘类型。
	VolumeType *string `json:"volume_type,omitempty"`

	// 共享存储大小，单位G。
	Size *int32 `json:"size,omitempty"`

	// 共享存储版本： - 0：共享存储1.0 - 1：共享存储2.0
	Version *int32 `json:"version,omitempty"`
}

func (o ShowCloudPhoneServerDetailResponseBodyShareVolumeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCloudPhoneServerDetailResponseBodyShareVolumeInfo struct{}"
	}

	return strings.Join([]string{"ShowCloudPhoneServerDetailResponseBodyShareVolumeInfo", string(data)}, " ")
}
