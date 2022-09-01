package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SourceInfo struct {

	// 事件源ID
	Id *string `json:"id,omitempty" xml:"id"`

	// 事件源名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 事件源名称展示
	Label *string `json:"label,omitempty" xml:"label"`

	// 事件源描述
	Description *string `json:"description,omitempty" xml:"description"`

	// 事件源提供方类型，OFFICIAL：官方云服务事件源；CUSTOM：用户创建的自定义事件源
	ProviderType *SourceInfoProviderType `json:"provider_type,omitempty" xml:"provider_type"`

	// 事件源提供的事件类型列表，只有官方云服务事件源提供事件类型
	EventTypes *[]SourceInfoEventTypes `json:"event_types,omitempty" xml:"event_types"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty" xml:"created_time"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty" xml:"updated_time"`

	// 事件源归属的事件通道ID
	ChannelId *string `json:"channel_id,omitempty" xml:"channel_id"`

	// 事件源归属的事件通道名称
	ChannelName *string `json:"channel_name,omitempty" xml:"channel_name"`
}

func (o SourceInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SourceInfo struct{}"
	}

	return strings.Join([]string{"SourceInfo", string(data)}, " ")
}

type SourceInfoProviderType struct {
	value string
}

type SourceInfoProviderTypeEnum struct {
	OFFICIAL SourceInfoProviderType
	CUSTOM   SourceInfoProviderType
}

func GetSourceInfoProviderTypeEnum() SourceInfoProviderTypeEnum {
	return SourceInfoProviderTypeEnum{
		OFFICIAL: SourceInfoProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: SourceInfoProviderType{
			value: "CUSTOM",
		},
	}
}

func (c SourceInfoProviderType) Value() string {
	return c.value
}

func (c SourceInfoProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SourceInfoProviderType) UnmarshalJSON(b []byte) error {
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
