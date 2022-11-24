package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 目的端磁盘
type TargetDisk struct {

	// 磁盘标识ID
	Id *int64 `json:"id,omitempty"`

	// 判断是普通分区，启动分区还是系统分区 BOOT：BOOT设备 OS：系统设备 NORMAL:平常
	DeviceUse *TargetDiskDeviceUse `json:"device_use,omitempty"`

	// 磁盘ID
	DiskId *string `json:"disk_id,omitempty"`

	// 磁盘名称
	Name *string `json:"name,omitempty"`

	// 逻辑卷信息
	PhysicalVolumes *[]TargetPhysicalVolumes `json:"physical_volumes,omitempty"`

	// 大小
	Size *int64 `json:"size,omitempty"`

	// 已使用大小
	UsedSize *int64 `json:"used_size,omitempty"`

	// 磁盘索引
	DiskIndex *string `json:"disk_index,omitempty"`

	// 是否为系统盘
	OsDisk *bool `json:"os_disk,omitempty"`

	// 磁盘的分区类型，添加源端时源端磁盘必选 MBR：mbr格式 GPT：gpt格式
	PartitionStyle *TargetDiskPartitionStyle `json:"partition_style,omitempty"`

	// Linux系统 目的端ECS中与源端关联的磁盘名称
	RelationName *string `json:"relation_name,omitempty"`
}

func (o TargetDisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetDisk struct{}"
	}

	return strings.Join([]string{"TargetDisk", string(data)}, " ")
}

type TargetDiskDeviceUse struct {
	value string
}

type TargetDiskDeviceUseEnum struct {
	NORMAL TargetDiskDeviceUse
	OS     TargetDiskDeviceUse
	BOOT   TargetDiskDeviceUse
}

func GetTargetDiskDeviceUseEnum() TargetDiskDeviceUseEnum {
	return TargetDiskDeviceUseEnum{
		NORMAL: TargetDiskDeviceUse{
			value: "NORMAL",
		},
		OS: TargetDiskDeviceUse{
			value: "OS",
		},
		BOOT: TargetDiskDeviceUse{
			value: "BOOT",
		},
	}
}

func (c TargetDiskDeviceUse) Value() string {
	return c.value
}

func (c TargetDiskDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetDiskDeviceUse) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type TargetDiskPartitionStyle struct {
	value string
}

type TargetDiskPartitionStyleEnum struct {
	MBR TargetDiskPartitionStyle
	GPT TargetDiskPartitionStyle
}

func GetTargetDiskPartitionStyleEnum() TargetDiskPartitionStyleEnum {
	return TargetDiskPartitionStyleEnum{
		MBR: TargetDiskPartitionStyle{
			value: "MBR",
		},
		GPT: TargetDiskPartitionStyle{
			value: "GPT",
		},
	}
}

func (c TargetDiskPartitionStyle) Value() string {
	return c.value
}

func (c TargetDiskPartitionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetDiskPartitionStyle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
