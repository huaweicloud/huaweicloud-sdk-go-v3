package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VolumeSpec struct {

	// 资源类型，当前只支持“obs”。
	ResourceType string `json:"resource_type"`

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
