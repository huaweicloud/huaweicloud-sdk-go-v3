package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateTagRequest Request Object
type CreateTagRequest struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 资源类型: - cloud-connection: 云连接 - bandwidth-package: 带宽包
	ResourceType CreateTagRequestResourceType `json:"resource_type"`

	Body *CreateTagRequestBody `json:"body,omitempty"`
}

func (o CreateTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagRequest struct{}"
	}

	return strings.Join([]string{"CreateTagRequest", string(data)}, " ")
}

type CreateTagRequestResourceType struct {
	value string
}

type CreateTagRequestResourceTypeEnum struct {
	CLOUD_CONNECTION  CreateTagRequestResourceType
	BANDWIDTH_PACKAGE CreateTagRequestResourceType
}

func GetCreateTagRequestResourceTypeEnum() CreateTagRequestResourceTypeEnum {
	return CreateTagRequestResourceTypeEnum{
		CLOUD_CONNECTION: CreateTagRequestResourceType{
			value: "cloud-connection",
		},
		BANDWIDTH_PACKAGE: CreateTagRequestResourceType{
			value: "bandwidth-package",
		},
	}
}

func (c CreateTagRequestResourceType) Value() string {
	return c.value
}

func (c CreateTagRequestResourceType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTagRequestResourceType) UnmarshalJSON(b []byte) error {
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
