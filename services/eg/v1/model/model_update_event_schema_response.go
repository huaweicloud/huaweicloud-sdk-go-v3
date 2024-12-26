package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateEventSchemaResponse Response Object
type UpdateEventSchemaResponse struct {

	// 事件模型ID
	Id *string `json:"id,omitempty"`

	// 事件模型名称，租户下唯一
	Name *string `json:"name,omitempty"`

	// 事件模型描述
	Description *string `json:"description,omitempty"`

	// 事件模型兼容性
	Compatibility *UpdateEventSchemaResponseCompatibility `json:"compatibility,omitempty"`

	// 提供方类型，OFFICIAL：官方事件源；CUSTOM：自定义事件源
	ProviderType *UpdateEventSchemaResponseProviderType `json:"provider_type,omitempty"`

	// 事件模型格式
	Format *string `json:"format,omitempty"`

	// 事件模型版本数
	NumberOfVersions *int32 `json:"number_of_versions,omitempty"`

	// 创建时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 事件模型当前版本号
	Version *int32 `json:"version,omitempty"`

	// 事件模型内容定义
	Definition *string `json:"definition,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateEventSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEventSchemaResponse struct{}"
	}

	return strings.Join([]string{"UpdateEventSchemaResponse", string(data)}, " ")
}

type UpdateEventSchemaResponseCompatibility struct {
	value string
}

type UpdateEventSchemaResponseCompatibilityEnum struct {
	NONE UpdateEventSchemaResponseCompatibility
}

func GetUpdateEventSchemaResponseCompatibilityEnum() UpdateEventSchemaResponseCompatibilityEnum {
	return UpdateEventSchemaResponseCompatibilityEnum{
		NONE: UpdateEventSchemaResponseCompatibility{
			value: "NONE",
		},
	}
}

func (c UpdateEventSchemaResponseCompatibility) Value() string {
	return c.value
}

func (c UpdateEventSchemaResponseCompatibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEventSchemaResponseCompatibility) UnmarshalJSON(b []byte) error {
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

type UpdateEventSchemaResponseProviderType struct {
	value string
}

type UpdateEventSchemaResponseProviderTypeEnum struct {
	OFFICIAL UpdateEventSchemaResponseProviderType
	CUSTOM   UpdateEventSchemaResponseProviderType
}

func GetUpdateEventSchemaResponseProviderTypeEnum() UpdateEventSchemaResponseProviderTypeEnum {
	return UpdateEventSchemaResponseProviderTypeEnum{
		OFFICIAL: UpdateEventSchemaResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: UpdateEventSchemaResponseProviderType{
			value: "CUSTOM",
		},
	}
}

func (c UpdateEventSchemaResponseProviderType) Value() string {
	return c.value
}

func (c UpdateEventSchemaResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateEventSchemaResponseProviderType) UnmarshalJSON(b []byte) error {
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
