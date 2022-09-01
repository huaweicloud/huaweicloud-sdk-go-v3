package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 物理分区
type TargetPhysicalVolumes struct {

	// 分区类型
	DeviceUse *TargetPhysicalVolumesDeviceUse `json:"device_use,omitempty" xml:"device_use"`

	// 文件系统
	FileSystem *string `json:"file_system,omitempty" xml:"file_system"`

	// 编号
	Index *int32 `json:"index,omitempty" xml:"index"`

	// 挂载点
	MountPoint *string `json:"mount_point,omitempty" xml:"mount_point"`

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 大小
	Size *int64 `json:"size,omitempty" xml:"size"`

	// 使用大小
	UsedSize *int64 `json:"used_size,omitempty" xml:"used_size"`

	// uuid
	Uuid *string `json:"uuid,omitempty" xml:"uuid"`
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
