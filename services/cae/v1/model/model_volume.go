package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Volume struct {

	// 云存储ID。
	Id *string `json:"id,omitempty"`

	// 存储资源详情。
	ResourceInfo map[string]string `json:"resource_info,omitempty"`

	// 并行文件系统或存储桶名称。
	ResourceName *string `json:"resource_name,omitempty"`

	// 资源类型，当前只支持“obs”。
	ResourceType *string `json:"resource_type,omitempty"`

	// 存储资源子类型，例如：并行文件系统、存储桶。
	ResourceSubType *VolumeResourceSubType `json:"resource_sub_type,omitempty"`

	// 创建时间。
	Time *string `json:"time,omitempty"`
}

func (o Volume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Volume struct{}"
	}

	return strings.Join([]string{"Volume", string(data)}, " ")
}

type VolumeResourceSubType struct {
	value string
}

type VolumeResourceSubTypeEnum struct {
	PARALLEL_FILE_SYSTEM VolumeResourceSubType
	OBJECT_BUCKET        VolumeResourceSubType
}

func GetVolumeResourceSubTypeEnum() VolumeResourceSubTypeEnum {
	return VolumeResourceSubTypeEnum{
		PARALLEL_FILE_SYSTEM: VolumeResourceSubType{
			value: "parallel_file_system",
		},
		OBJECT_BUCKET: VolumeResourceSubType{
			value: "object_bucket",
		},
	}
}

func (c VolumeResourceSubType) Value() string {
	return c.value
}

func (c VolumeResourceSubType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *VolumeResourceSubType) UnmarshalJSON(b []byte) error {
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
