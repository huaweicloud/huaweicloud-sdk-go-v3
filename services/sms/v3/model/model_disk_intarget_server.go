package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DiskIntargetServer 目的端服务器关联磁盘
type DiskIntargetServer struct {

	// 磁盘名称
	Name string `json:"name"`

	// 磁盘大小，单位：字节
	Size int64 `json:"size"`

	// 磁盘的作用 BOOT：BOOT设备 OS：系统设备 NORMAL:平常
	DeviceUse *DiskIntargetServerDeviceUse `json:"device_use,omitempty"`

	// 磁盘已使用大小，以字节为单位
	UsedSize *int64 `json:"used_size,omitempty"`

	// 物理卷信息
	PhysicalVolumes *[]PhysicalVolumes `json:"physical_volumes,omitempty"`
}

func (o DiskIntargetServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DiskIntargetServer struct{}"
	}

	return strings.Join([]string{"DiskIntargetServer", string(data)}, " ")
}

type DiskIntargetServerDeviceUse struct {
	value string
}

type DiskIntargetServerDeviceUseEnum struct {
	BOOT   DiskIntargetServerDeviceUse
	OS     DiskIntargetServerDeviceUse
	NORMAL DiskIntargetServerDeviceUse
}

func GetDiskIntargetServerDeviceUseEnum() DiskIntargetServerDeviceUseEnum {
	return DiskIntargetServerDeviceUseEnum{
		BOOT: DiskIntargetServerDeviceUse{
			value: "BOOT",
		},
		OS: DiskIntargetServerDeviceUse{
			value: "OS",
		},
		NORMAL: DiskIntargetServerDeviceUse{
			value: "NORMAL",
		},
	}
}

func (c DiskIntargetServerDeviceUse) Value() string {
	return c.value
}

func (c DiskIntargetServerDeviceUse) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DiskIntargetServerDeviceUse) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
