package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ProductTopic struct {

	// 归属产品ID
	ProductId *int32 `json:"product_id,omitempty" xml:"product_id"`

	// 产品主题ID
	TopicId *string `json:"topic_id,omitempty" xml:"topic_id"`

	// 主题权限 0-发布 1-订阅
	Permission *ProductTopicPermission `json:"permission,omitempty" xml:"permission"`

	// 主题名称
	TopicName *string `json:"topic_name,omitempty" xml:"topic_name"`

	// 版本号
	Version *string `json:"version,omitempty" xml:"version"`

	// 描述
	Description *string `json:"description,omitempty" xml:"description"`
}

func (o ProductTopic) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductTopic struct{}"
	}

	return strings.Join([]string{"ProductTopic", string(data)}, " ")
}

type ProductTopicPermission struct {
	value int32
}

type ProductTopicPermissionEnum struct {
	E_0 ProductTopicPermission
	E_1 ProductTopicPermission
}

func GetProductTopicPermissionEnum() ProductTopicPermissionEnum {
	return ProductTopicPermissionEnum{
		E_0: ProductTopicPermission{
			value: 0,
		}, E_1: ProductTopicPermission{
			value: 1,
		},
	}
}

func (c ProductTopicPermission) Value() int32 {
	return c.value
}

func (c ProductTopicPermission) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProductTopicPermission) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
