package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ProductTopic struct {
	// 归属产品ID

	ProductId *int32 `json:"product_id,omitempty"`
	// 产品主题ID

	TopicId *string `json:"topic_id,omitempty"`
	// 主题权限 0-发布 1-订阅

	Permission *ProductTopicPermission `json:"permission,omitempty"`
	// 主题名称

	TopicName *string `json:"topic_name,omitempty"`
	// 版本号

	Version *string `json:"version,omitempty"`
	// 描述

	Description *string `json:"description,omitempty"`
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
