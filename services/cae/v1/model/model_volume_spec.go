package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VolumeSpec struct {

	// 资源类型。
	ResourceType VolumeSpecResourceType `json:"resource_type"`

	// 对象存储类型，例如：并行文件系统、存储桶。
	ResourceSubType VolumeSpecResourceSubType `json:"resource_sub_type"`

	// 并行文件系统或存储桶名称。
	Resources []string `json:"resources"`
}

func (o VolumeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeSpec struct{}"
	}

	return strings.Join([]string{"VolumeSpec", string(data)}, " ")
}

type VolumeSpecResourceType struct {
	value string
}

type VolumeSpecResourceTypeEnum struct {
	OBS VolumeSpecResourceType
}

func GetVolumeSpecResourceTypeEnum() VolumeSpecResourceTypeEnum {
	return VolumeSpecResourceTypeEnum{
		OBS: VolumeSpecResourceType{
			value: "obs",
		},
	}
}

func (c VolumeSpecResourceType) Value() string {
	return c.value
}

func (c VolumeSpecResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeSpecResourceType) UnmarshalJSON(b []byte) error {
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

type VolumeSpecResourceSubType struct {
	value string
}

type VolumeSpecResourceSubTypeEnum struct {
	OBJECT_BUCKET        VolumeSpecResourceSubType
	PARALLEL_FILE_SYSTEM VolumeSpecResourceSubType
}

func GetVolumeSpecResourceSubTypeEnum() VolumeSpecResourceSubTypeEnum {
	return VolumeSpecResourceSubTypeEnum{
		OBJECT_BUCKET: VolumeSpecResourceSubType{
			value: "object_bucket",
		},
		PARALLEL_FILE_SYSTEM: VolumeSpecResourceSubType{
			value: "parallel_file_system",
		},
	}
}

func (c VolumeSpecResourceSubType) Value() string {
	return c.value
}

func (c VolumeSpecResourceSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeSpecResourceSubType) UnmarshalJSON(b []byte) error {
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
