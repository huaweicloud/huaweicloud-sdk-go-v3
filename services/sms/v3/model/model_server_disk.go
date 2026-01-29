package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerDisk 磁盘信息
type ServerDisk struct {

	// 磁盘名称
	Name *string `json:"name,omitempty"`

	// 磁盘的分区类型，添加源端时源端磁盘必选，否则无法通过后续环境检查 （非枚举数据，来源于EVS服务） 常见类型如：MBR：主启动记录分区，GPT：Guid Partition Table，全局分区表。 详细类型请参考EVS服务API文档中“MBR和GPT分区形式有何区别”部分描述
	PartitionStyle *string `json:"partition_style,omitempty"`

	// 磁盘类型。 无强约束，可为空值，常见取值如下 NORMAL：平常 OS：系统设备 BOOT：BOOT设备 VOLUME_GROUP：VolumeGroup组成设备 BTRFS：BTRFS组成设备
	DeviceUse *string `json:"device_use,omitempty"`

	// 磁盘总大小，以字节为单位
	Size *int64 `json:"size,omitempty"`

	// 磁盘已使用大小，以字节为单位
	UsedSize *int64 `json:"used_size,omitempty"`

	// 磁盘上的物理分区信息
	PhysicalVolumes *[]PhysicalVolume `json:"physical_volumes,omitempty"`

	// 是否为系统盘
	OsDisk *bool `json:"os_disk,omitempty"`

	// Linux系统 目的端ECS中与源端关联的磁盘名称
	RelationName *string `json:"relation_name,omitempty"`

	// inode大小
	InodeSize *int32 `json:"inode_size,omitempty"`
}

func (o ServerDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerDisk struct{}"
	}

	return strings.Join([]string{"ServerDisk", string(data)}, " ")
}
