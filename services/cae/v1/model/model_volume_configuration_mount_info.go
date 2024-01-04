package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeConfigurationMountInfo 云存储配置挂载信息。
type VolumeConfigurationMountInfo struct {

	// 容器挂载路径。
	Path *string `json:"path,omitempty"`

	// 子路径。
	SubPath *string `json:"sub_path,omitempty"`

	// 读写权限。
	AccessMode *VolumeConfigurationMountInfoAccessMode `json:"access_mode,omitempty"`
}

func (o VolumeConfigurationMountInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeConfigurationMountInfo struct{}"
	}

	return strings.Join([]string{"VolumeConfigurationMountInfo", string(data)}, " ")
}

type VolumeConfigurationMountInfoAccessMode struct {
	value string
}

type VolumeConfigurationMountInfoAccessModeEnum struct {
	READ_WRITE_MANY VolumeConfigurationMountInfoAccessMode
	READ_ONLY_MANY  VolumeConfigurationMountInfoAccessMode
}

func GetVolumeConfigurationMountInfoAccessModeEnum() VolumeConfigurationMountInfoAccessModeEnum {
	return VolumeConfigurationMountInfoAccessModeEnum{
		READ_WRITE_MANY: VolumeConfigurationMountInfoAccessMode{
			value: "ReadWriteMany",
		},
		READ_ONLY_MANY: VolumeConfigurationMountInfoAccessMode{
			value: "ReadOnlyMany",
		},
	}
}

func (c VolumeConfigurationMountInfoAccessMode) Value() string {
	return c.value
}

func (c VolumeConfigurationMountInfoAccessMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeConfigurationMountInfoAccessMode) UnmarshalJSON(b []byte) error {
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
