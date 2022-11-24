package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 请求数据。
type CreateEnvironmentRequestBodyMetadata struct {

	// 资源信息。
	Annotations map[string]string `json:"annotations,omitempty"`

	// 环境名称。
	Name string `json:"name"`

	// 环境类型。
	Type CreateEnvironmentRequestBodyMetadataType `json:"type"`
}

func (o CreateEnvironmentRequestBodyMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEnvironmentRequestBodyMetadata struct{}"
	}

	return strings.Join([]string{"CreateEnvironmentRequestBodyMetadata", string(data)}, " ")
}

type CreateEnvironmentRequestBodyMetadataType struct {
	value string
}

type CreateEnvironmentRequestBodyMetadataTypeEnum struct {
	EXCLUSIVE CreateEnvironmentRequestBodyMetadataType
}

func GetCreateEnvironmentRequestBodyMetadataTypeEnum() CreateEnvironmentRequestBodyMetadataTypeEnum {
	return CreateEnvironmentRequestBodyMetadataTypeEnum{
		EXCLUSIVE: CreateEnvironmentRequestBodyMetadataType{
			value: "exclusive",
		},
	}
}

func (c CreateEnvironmentRequestBodyMetadataType) Value() string {
	return c.value
}

func (c CreateEnvironmentRequestBodyMetadataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEnvironmentRequestBodyMetadataType) UnmarshalJSON(b []byte) error {
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
