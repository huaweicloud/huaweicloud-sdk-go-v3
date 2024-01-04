package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// VolumeConfigurationDataVolume 云存储配置数据。
type VolumeConfigurationDataVolume struct {

	// 云存储ID。
	VolumeId *string `json:"volume_id,omitempty"`

	// 云存储名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 云存储类型。
	ResourceType *VolumeConfigurationDataVolumeResourceType `json:"resource_type,omitempty"`

	// 云存储子类型。
	ResourceSubType *VolumeConfigurationDataVolumeResourceSubType `json:"resource_sub_type,omitempty"`

	MountInfo *[]VolumeConfigurationMountInfo `json:"mount_info,omitempty"`
}

func (o VolumeConfigurationDataVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeConfigurationDataVolume struct{}"
	}

	return strings.Join([]string{"VolumeConfigurationDataVolume", string(data)}, " ")
}

type VolumeConfigurationDataVolumeResourceType struct {
	value string
}

type VolumeConfigurationDataVolumeResourceTypeEnum struct {
	OBS VolumeConfigurationDataVolumeResourceType
	SFS VolumeConfigurationDataVolumeResourceType
}

func GetVolumeConfigurationDataVolumeResourceTypeEnum() VolumeConfigurationDataVolumeResourceTypeEnum {
	return VolumeConfigurationDataVolumeResourceTypeEnum{
		OBS: VolumeConfigurationDataVolumeResourceType{
			value: "obs",
		},
		SFS: VolumeConfigurationDataVolumeResourceType{
			value: "sfs",
		},
	}
}

func (c VolumeConfigurationDataVolumeResourceType) Value() string {
	return c.value
}

func (c VolumeConfigurationDataVolumeResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeConfigurationDataVolumeResourceType) UnmarshalJSON(b []byte) error {
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

type VolumeConfigurationDataVolumeResourceSubType struct {
	value string
}

type VolumeConfigurationDataVolumeResourceSubTypeEnum struct {
	OBJECT_BUCKET        VolumeConfigurationDataVolumeResourceSubType
	PARALLEL_FILE_SYSTEM VolumeConfigurationDataVolumeResourceSubType
	SFS3_0               VolumeConfigurationDataVolumeResourceSubType
}

func GetVolumeConfigurationDataVolumeResourceSubTypeEnum() VolumeConfigurationDataVolumeResourceSubTypeEnum {
	return VolumeConfigurationDataVolumeResourceSubTypeEnum{
		OBJECT_BUCKET: VolumeConfigurationDataVolumeResourceSubType{
			value: "object_bucket",
		},
		PARALLEL_FILE_SYSTEM: VolumeConfigurationDataVolumeResourceSubType{
			value: "parallel_file_system",
		},
		SFS3_0: VolumeConfigurationDataVolumeResourceSubType{
			value: "sfs3.0",
		},
	}
}

func (c VolumeConfigurationDataVolumeResourceSubType) Value() string {
	return c.value
}

func (c VolumeConfigurationDataVolumeResourceSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeConfigurationDataVolumeResourceSubType) UnmarshalJSON(b []byte) error {
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
