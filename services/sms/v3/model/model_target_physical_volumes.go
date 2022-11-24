package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 物理分区
type TargetPhysicalVolumes struct {

	// 逻辑卷ID
	Id *int64 `json:"id,omitempty"`

	// 分区类型 NORMAL:平常 OS：系统设备 BOOT：BOOT设备
	DeviceUse *TargetPhysicalVolumesDeviceUse `json:"device_use,omitempty"`

	// 文件系统
	FileSystem *string `json:"file_system,omitempty"`

	// 编号
	Index *int32 `json:"index,omitempty"`

	// 挂载点
	MountPoint *string `json:"mount_point,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 大小
	Size *int64 `json:"size,omitempty"`

	// 使用大小
	UsedSize *int64 `json:"used_size,omitempty"`

	// uuid
	Uuid *string `json:"uuid,omitempty"`

	// Linux系统 目的端ECS中与源端关联的磁盘名称
	RelationName *string `json:"relation_name,omitempty"`

	// 分区空闲大小
	FreeSize *int64 `json:"free_size,omitempty"`
}

func (o TargetPhysicalVolumes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TargetPhysicalVolumes struct{}"
	}

	return strings.Join([]string{"TargetPhysicalVolumes", string(data)}, " ")
}

type TargetPhysicalVolumesDeviceUse struct {
	value string
}

type TargetPhysicalVolumesDeviceUseEnum struct {
	NORMAL TargetPhysicalVolumesDeviceUse
	OS     TargetPhysicalVolumesDeviceUse
	BOOT   TargetPhysicalVolumesDeviceUse
}

func GetTargetPhysicalVolumesDeviceUseEnum() TargetPhysicalVolumesDeviceUseEnum {
	return TargetPhysicalVolumesDeviceUseEnum{
		NORMAL: TargetPhysicalVolumesDeviceUse{
			value: "NORMAL",
		},
		OS: TargetPhysicalVolumesDeviceUse{
			value: "OS",
		},
		BOOT: TargetPhysicalVolumesDeviceUse{
			value: "BOOT",
		},
	}
}

func (c TargetPhysicalVolumesDeviceUse) Value() string {
	return c.value
}

func (c TargetPhysicalVolumesDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TargetPhysicalVolumesDeviceUse) UnmarshalJSON(b []byte) error {
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
