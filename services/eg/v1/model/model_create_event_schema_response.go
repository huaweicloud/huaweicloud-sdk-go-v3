package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateEventSchemaResponse struct {

	// 事件模型ID
	Id *string `json:"id,omitempty"`

	// 事件模型名称，租户下唯一
	Name *string `json:"name,omitempty"`

	// 事件模型描述
	Description *string `json:"description,omitempty"`

	// 事件模型兼容性
	Compatibility *CreateEventSchemaResponseCompatibility `json:"compatibility,omitempty"`

	// 提供方类型，OFFICIAL：官方事件源；CUSTOM：自定义事件源
	ProviderType *CreateEventSchemaResponseProviderType `json:"provider_type,omitempty"`

	// 事件模型格式
	Format *string `json:"format,omitempty"`

	// 事件模型版本数
	NumberOfVersions *int32 `json:"number_of_versions,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 事件示例
	DataSample *string `json:"data_sample,omitempty"`

	// 事件模型当前版本号
	Version *int32 `json:"version,omitempty"`

	// 事件模型内容定义
	Definition *string `json:"definition,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSchemaResponse struct{}"
	}

	return strings.Join([]string{"CreateEventSchemaResponse", string(data)}, " ")
}

type CreateEventSchemaResponseCompatibility struct {
	value string
}

type CreateEventSchemaResponseCompatibilityEnum struct {
	NONE     CreateEventSchemaResponseCompatibility
	FORWARD  CreateEventSchemaResponseCompatibility
	BACKWARD CreateEventSchemaResponseCompatibility
	FULL     CreateEventSchemaResponseCompatibility
}

func GetCreateEventSchemaResponseCompatibilityEnum() CreateEventSchemaResponseCompatibilityEnum {
	return CreateEventSchemaResponseCompatibilityEnum{
		NONE: CreateEventSchemaResponseCompatibility{
			value: "NONE",
		},
		FORWARD: CreateEventSchemaResponseCompatibility{
			value: "FORWARD",
		},
		BACKWARD: CreateEventSchemaResponseCompatibility{
			value: "BACKWARD",
		},
		FULL: CreateEventSchemaResponseCompatibility{
			value: "FULL",
		},
	}
}

func (c CreateEventSchemaResponseCompatibility) Value() string {
	return c.value
}

func (c CreateEventSchemaResponseCompatibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEventSchemaResponseCompatibility) UnmarshalJSON(b []byte) error {
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

type CreateEventSchemaResponseProviderType struct {
	value string
}

type CreateEventSchemaResponseProviderTypeEnum struct {
	OFFICIAL CreateEventSchemaResponseProviderType
	CUSTOM   CreateEventSchemaResponseProviderType
}

func GetCreateEventSchemaResponseProviderTypeEnum() CreateEventSchemaResponseProviderTypeEnum {
	return CreateEventSchemaResponseProviderTypeEnum{
		OFFICIAL: CreateEventSchemaResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CreateEventSchemaResponseProviderType{
			value: "CUSTOM",
		},
	}
}

func (c CreateEventSchemaResponseProviderType) Value() string {
	return c.value
}

func (c CreateEventSchemaResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEventSchemaResponseProviderType) UnmarshalJSON(b []byte) error {
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
