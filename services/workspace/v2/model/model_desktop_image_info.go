package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopImageInfo 桌面镜像信息响应。
type DesktopImageInfo struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 镜像名称。
	ImageName *string `json:"image_name,omitempty"`

	// 镜像类型。
	ImageType *string `json:"image_type,omitempty"`

	// 镜像系统类型。
	OsType *string `json:"os_type,omitempty"`
}

func (o DesktopImageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopImageInfo struct{}"
	}

	return strings.Join([]string{"DesktopImageInfo", string(data)}, " ")
}
