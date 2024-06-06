package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type VolumeSpec struct {

	// 资源类型，当前只支持“obs”和“sfs”。
	ResourceType string `json:"resource_type"`

	// 对象存储类型。
	ResourceSubType VolumeSpecResourceSubType `json:"resource_sub_type"`

	// 云存储名称。
	Resources []string `json:"resources"`

	// 云存储和授权凭证，获取环境列表接口响应中env_category字段为v2时需添加该字段的值。
	ResourcesCredentials *[]ResourcesCredential `json:"resources_credentials,omitempty"`
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
	SFS3_0               VolumeSpecResourceSubType
}

func GetVolumeSpecResourceSubTypeEnum() VolumeSpecResourceSubTypeEnum {
	return VolumeSpecResourceSubTypeEnum{
		OBJECT_BUCKET: VolumeSpecResourceSubType{
			value: "object_bucket",
		},
		PARALLEL_FILE_SYSTEM: VolumeSpecResourceSubType{
			value: "parallel_file_system",
		},
		SFS3_0: VolumeSpecResourceSubType{
			value: "sfs3.0",
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
