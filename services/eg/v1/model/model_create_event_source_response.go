package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateEventSourceResponse struct {

	// 事件源ID
	Id *string `json:"id,omitempty"`

	// 事件源名称
	Name *string `json:"name,omitempty"`

	// 事件源名称展示
	Label *string `json:"label,omitempty"`

	// 事件源描述
	Description *string `json:"description,omitempty"`

	// 事件源提供方类型，OFFICIAL：官方云服务事件源；CUSTOM：用户创建的自定义事件源
	ProviderType *CreateEventSourceResponseProviderType `json:"provider_type,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	// 事件源归属的事件通道ID
	ChannelId *string `json:"channel_id,omitempty"`

	// 事件源归属的事件通道名称
	ChannelName    *string `json:"channel_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateEventSourceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEventSourceResponse struct{}"
	}

	return strings.Join([]string{"CreateEventSourceResponse", string(data)}, " ")
}

type CreateEventSourceResponseProviderType struct {
	value string
}

type CreateEventSourceResponseProviderTypeEnum struct {
	OFFICIAL CreateEventSourceResponseProviderType
	CUSTOM   CreateEventSourceResponseProviderType
}

func GetCreateEventSourceResponseProviderTypeEnum() CreateEventSourceResponseProviderTypeEnum {
	return CreateEventSourceResponseProviderTypeEnum{
		OFFICIAL: CreateEventSourceResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CreateEventSourceResponseProviderType{
			value: "CUSTOM",
		},
	}
}

func (c CreateEventSourceResponseProviderType) Value() string {
	return c.value
}

func (c CreateEventSourceResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEventSourceResponseProviderType) UnmarshalJSON(b []byte) error {
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
