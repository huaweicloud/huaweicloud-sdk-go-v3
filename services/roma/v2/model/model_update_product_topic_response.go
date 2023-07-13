package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateProductTopicResponse Response Object
type UpdateProductTopicResponse struct {

	// 归属产品ID
	ProductId *int32 `json:"product_id,omitempty"`

	// 产品主题ID
	TopicId *int32 `json:"topic_id,omitempty"`

	// 主题权限 0-发布 1-订阅
	Permission *UpdateProductTopicResponsePermission `json:"permission,omitempty"`

	// 主题名称
	TopicName *string `json:"topic_name,omitempty"`

	// 版本号
	Version *string `json:"version,omitempty"`

	// 描述
	Description    *string `json:"description,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProductTopicResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProductTopicResponse struct{}"
	}

	return strings.Join([]string{"UpdateProductTopicResponse", string(data)}, " ")
}

type UpdateProductTopicResponsePermission struct {
	value int32
}

type UpdateProductTopicResponsePermissionEnum struct {
	E_0 UpdateProductTopicResponsePermission
	E_1 UpdateProductTopicResponsePermission
}

func GetUpdateProductTopicResponsePermissionEnum() UpdateProductTopicResponsePermissionEnum {
	return UpdateProductTopicResponsePermissionEnum{
		E_0: UpdateProductTopicResponsePermission{
			value: 0,
		}, E_1: UpdateProductTopicResponsePermission{
			value: 1,
		},
	}
}

func (c UpdateProductTopicResponsePermission) Value() int32 {
	return c.value
}

func (c UpdateProductTopicResponsePermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateProductTopicResponsePermission) UnmarshalJSON(b []byte) error {
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
