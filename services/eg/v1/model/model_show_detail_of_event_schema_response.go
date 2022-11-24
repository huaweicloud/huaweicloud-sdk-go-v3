package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowDetailOfEventSchemaResponse struct {

	// 事件模型ID
	Id *string `json:"id,omitempty"`

	// 事件模型名称，租户下唯一
	Name *string `json:"name,omitempty"`

	// 事件模型描述
	Description *string `json:"description,omitempty"`

	// 事件模型兼容性
	Compatibility *ShowDetailOfEventSchemaResponseCompatibility `json:"compatibility,omitempty"`

	// 提供方类型，OFFICIAL：官方事件源；CUSTOM：自定义事件源
	ProviderType *ShowDetailOfEventSchemaResponseProviderType `json:"provider_type,omitempty"`

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
	Definition     *string `json:"definition,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfEventSchemaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfEventSchemaResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfEventSchemaResponse", string(data)}, " ")
}

type ShowDetailOfEventSchemaResponseCompatibility struct {
	value string
}

type ShowDetailOfEventSchemaResponseCompatibilityEnum struct {
	NONE     ShowDetailOfEventSchemaResponseCompatibility
	FORWARD  ShowDetailOfEventSchemaResponseCompatibility
	BACKWARD ShowDetailOfEventSchemaResponseCompatibility
	FULL     ShowDetailOfEventSchemaResponseCompatibility
}

func GetShowDetailOfEventSchemaResponseCompatibilityEnum() ShowDetailOfEventSchemaResponseCompatibilityEnum {
	return ShowDetailOfEventSchemaResponseCompatibilityEnum{
		NONE: ShowDetailOfEventSchemaResponseCompatibility{
			value: "NONE",
		},
		FORWARD: ShowDetailOfEventSchemaResponseCompatibility{
			value: "FORWARD",
		},
		BACKWARD: ShowDetailOfEventSchemaResponseCompatibility{
			value: "BACKWARD",
		},
		FULL: ShowDetailOfEventSchemaResponseCompatibility{
			value: "FULL",
		},
	}
}

func (c ShowDetailOfEventSchemaResponseCompatibility) Value() string {
	return c.value
}

func (c ShowDetailOfEventSchemaResponseCompatibility) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfEventSchemaResponseCompatibility) UnmarshalJSON(b []byte) error {
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

type ShowDetailOfEventSchemaResponseProviderType struct {
	value string
}

type ShowDetailOfEventSchemaResponseProviderTypeEnum struct {
	OFFICIAL ShowDetailOfEventSchemaResponseProviderType
	CUSTOM   ShowDetailOfEventSchemaResponseProviderType
}

func GetShowDetailOfEventSchemaResponseProviderTypeEnum() ShowDetailOfEventSchemaResponseProviderTypeEnum {
	return ShowDetailOfEventSchemaResponseProviderTypeEnum{
		OFFICIAL: ShowDetailOfEventSchemaResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ShowDetailOfEventSchemaResponseProviderType{
			value: "CUSTOM",
		},
	}
}

func (c ShowDetailOfEventSchemaResponseProviderType) Value() string {
	return c.value
}

func (c ShowDetailOfEventSchemaResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfEventSchemaResponseProviderType) UnmarshalJSON(b []byte) error {
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
