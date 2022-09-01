package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 磁盘信息
type Disk struct {

	// 磁盘名称
	Name string `json:"name" xml:"name"`

	// 磁盘的分区类型，添加源端时源端磁盘必选
	PartitionStyle *DiskPartitionStyle `json:"partition_style,omitempty" xml:"partition_style"`

	// 磁盘类型
	DeviceUse DiskDeviceUse `json:"device_use" xml:"device_use"`

	// 磁盘总大小，以字节为单位
	Size int64 `json:"size" xml:"size"`

	// 磁盘已使用大小，以字节为单位
	UsedSize int64 `json:"used_size" xml:"used_size"`

	// 磁盘上的物理分区信息
	PhysicalVolumes []PhysicalVolumes `json:"physical_volumes" xml:"physical_volumes"`

	// 创建任务时，如果选择已有虚拟机，此参数必选
	DiskId *string `json:"disk_id,omitempty" xml:"disk_id"`

	// 是否为系统盘
	OsDisk *bool `json:"os_disk,omitempty" xml:"os_disk"`

	// Linux系统 目的端ECS中与源端关联的磁盘名称
	RelationName *string `json:"relation_name,omitempty" xml:"relation_name"`
}

func (o Disk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Disk struct{}"
	}

	return strings.Join([]string{"Disk", string(data)}, " ")
}

type DiskPartitionStyle struct {
	value string
}

type DiskPartitionStyleEnum struct {
	MBR DiskPartitionStyle
	GPT DiskPartitionStyle
}

func GetDiskPartitionStyleEnum() DiskPartitionStyleEnum {
	return DiskPartitionStyleEnum{
		MBR: DiskPartitionStyle{
			value: "MBR",
		},
		GPT: DiskPartitionStyle{
			value: "GPT",
		},
	}
}

func (c DiskPartitionStyle) Value() string {
	return c.value
}

func (c DiskPartitionStyle) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskPartitionStyle) UnmarshalJSON(b []byte) error {
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

type DiskDeviceUse struct {
	value string
}

type DiskDeviceUseEnum struct {
	BOOT DiskDeviceUse
	OS   DiskDeviceUse
}

func GetDiskDeviceUseEnum() DiskDeviceUseEnum {
	return DiskDeviceUseEnum{
		BOOT: DiskDeviceUse{
			value: "BOOT",
		},
		OS: DiskDeviceUse{
			value: "OS",
		},
	}
}

func (c DiskDeviceUse) Value() string {
	return c.value
}

func (c DiskDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskDeviceUse) UnmarshalJSON(b []byte) error {
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
