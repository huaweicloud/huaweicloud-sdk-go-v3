package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CustomizeSchemaItemInfo struct {

	// 事件模型ID
	Id *string `json:"id,omitempty"`

	// 事件模型名称，租户下唯一
	Name *string `json:"name,omitempty"`

	// 事件模型描述
	Description *string `json:"description,omitempty"`

	// 事件模型兼容性
	Compatibility *CustomizeSchemaItemInfoCompatibility `json:"compatibility,omitempty"`

	// 提供方类型，OFFICIAL：官方事件源；CUSTOM：自定义事件源
	ProviderType *CustomizeSchemaItemInfoProviderType `json:"provider_type,omitempty"`

	// 事件模型格式
	Format *string `json:"format,omitempty"`

	// 事件模型版本数
	NumberOfVersions *int32 `json:"number_of_versions,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`
}

func (o CustomizeSchemaItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomizeSchemaItemInfo struct{}"
	}

	return strings.Join([]string{"CustomizeSchemaItemInfo", string(data)}, " ")
}

type CustomizeSchemaItemInfoCompatibility struct {
	value string
}

type CustomizeSchemaItemInfoCompatibilityEnum struct {
	NONE     CustomizeSchemaItemInfoCompatibility
	FORWARD  CustomizeSchemaItemInfoCompatibility
	BACKWARD CustomizeSchemaItemInfoCompatibility
	FULL     CustomizeSchemaItemInfoCompatibility
}

func GetCustomizeSchemaItemInfoCompatibilityEnum() CustomizeSchemaItemInfoCompatibilityEnum {
	return CustomizeSchemaItemInfoCompatibilityEnum{
		NONE: CustomizeSchemaItemInfoCompatibility{
			value: "NONE",
		},
		FORWARD: CustomizeSchemaItemInfoCompatibility{
			value: "FORWARD",
		},
		BACKWARD: CustomizeSchemaItemInfoCompatibility{
			value: "BACKWARD",
		},
		FULL: CustomizeSchemaItemInfoCompatibility{
			value: "FULL",
		},
	}
}

func (c CustomizeSchemaItemInfoCompatibility) Value() string {
	return c.value
}

func (c CustomizeSchemaItemInfoCompatibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSchemaItemInfoCompatibility) UnmarshalJSON(b []byte) error {
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

type CustomizeSchemaItemInfoProviderType struct {
	value string
}

type CustomizeSchemaItemInfoProviderTypeEnum struct {
	OFFICIAL CustomizeSchemaItemInfoProviderType
	CUSTOM   CustomizeSchemaItemInfoProviderType
}

func GetCustomizeSchemaItemInfoProviderTypeEnum() CustomizeSchemaItemInfoProviderTypeEnum {
	return CustomizeSchemaItemInfoProviderTypeEnum{
		OFFICIAL: CustomizeSchemaItemInfoProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CustomizeSchemaItemInfoProviderType{
			value: "CUSTOM",
		},
	}
}

func (c CustomizeSchemaItemInfoProviderType) Value() string {
	return c.value
}

func (c CustomizeSchemaItemInfoProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CustomizeSchemaItemInfoProviderType) UnmarshalJSON(b []byte) error {
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
