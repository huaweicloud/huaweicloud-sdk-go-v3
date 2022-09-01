package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 使用大小
type PhysicalVolume struct {

	// 分区类型，普通分区，启动分区，系统分区
	DeviceUse *string `json:"device_use,omitempty" xml:"device_use"`

	// 文件系统类型
	FileSystem *string `json:"file_system,omitempty" xml:"file_system"`

	// 顺序
	Index *int32 `json:"index,omitempty" xml:"index"`

	// 挂载点
	MountPoint *string `json:"mount_point,omitempty" xml:"mount_point"`

	// 名称，windows表示盘符，Linux表示设备号
	Name *string `json:"name,omitempty" xml:"name"`

	// 大小
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 使用大小
	UsedSize *int64 `json:"used_size,omitempty" xml:"used_size"`

	// GUID，可从源端查询
	Uuid *string `json:"uuid,omitempty" xml:"uuid"`

	// 每个cluster大小
	SizePerCluster *int32 `json:"size_per_cluster,omitempty" xml:"size_per_cluster"`
}

func (o PhysicalVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PhysicalVolume struct{}"
	}

	return strings.Join([]string{"PhysicalVolume", string(data)}, " ")
}
