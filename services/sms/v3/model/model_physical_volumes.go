package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 使用大小
type PhysicalVolumes struct {
	// 分区类型，普通分区，启动分区，系统分区

	DeviceUse *string `json:"device_use,omitempty"`
	// 文件系统类型

	FileSystem *string `json:"file_system,omitempty"`
	// 顺序

	Index *int32 `json:"index,omitempty"`
	// 挂载点

	MountPoint *string `json:"mount_point,omitempty"`
	// 名称，windows表示盘符，Linux表示设备号

	Name *string `json:"name,omitempty"`
	// 大小

	Size *int64 `json:"size,omitempty"`
	// 使用大小

	UsedSize *int64 `json:"used_size,omitempty"`
	// GUID，可从源端查询

	Uuid *string `json:"uuid,omitempty"`
}

func (o PhysicalVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhysicalVolumes struct{}"
	}

	return strings.Join([]string{"PhysicalVolumes", string(data)}, " ")
}
