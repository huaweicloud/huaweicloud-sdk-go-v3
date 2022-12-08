package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type CreateTagRequest struct {

	// 资源ID。
	ResourceId string `json:"resource_id"`

	// 资源类型: - cc: 云连接 - bwp: 带宽包
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
	CC  CreateTagRequestResourceType
	BWP CreateTagRequestResourceType
}

func GetCreateTagRequestResourceTypeEnum() CreateTagRequestResourceTypeEnum {
	return CreateTagRequestResourceTypeEnum{
		CC: CreateTagRequestResourceType{
			value: "cc",
		},
		BWP: CreateTagRequestResourceType{
			value: "bwp",
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
