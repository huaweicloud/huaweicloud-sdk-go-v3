package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateProductTopicResponse Response Object
type CreateProductTopicResponse struct {

	// 归属产品ID
	ProductId *int32 `json:"product_id,omitempty"`

	// 产品主题ID
	TopicId *string `json:"topic_id,omitempty"`

	// 主题权限 0-发布 1-订阅
	Permission *CreateProductTopicResponsePermission `json:"permission,omitempty"`

	// 主题名称
	TopicName *string `json:"topic_name,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProductTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProductTopicResponse struct{}"
	}

	return strings.Join([]string{"CreateProductTopicResponse", string(data)}, " ")
}

type CreateProductTopicResponsePermission struct {
	value int32
}

type CreateProductTopicResponsePermissionEnum struct {
	E_0 CreateProductTopicResponsePermission
	E_1 CreateProductTopicResponsePermission
}

func GetCreateProductTopicResponsePermissionEnum() CreateProductTopicResponsePermissionEnum {
	return CreateProductTopicResponsePermissionEnum{
		E_0: CreateProductTopicResponsePermission{
			value: 0,
		}, E_1: CreateProductTopicResponsePermission{
			value: 1,
		},
	}
}

func (c CreateProductTopicResponsePermission) Value() int32 {
	return c.value
}

func (c CreateProductTopicResponsePermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateProductTopicResponsePermission) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
