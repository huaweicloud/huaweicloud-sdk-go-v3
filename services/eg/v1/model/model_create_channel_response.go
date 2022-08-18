package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateChannelResponse struct {

	// 通道ID
	Id *string `json:"id,omitempty"`

	// 通道名称
	Name *string `json:"name,omitempty"`

	// 通道描述
	Description *string `json:"description,omitempty"`

	// 通道提供方类型，OFFICIAL：官方事件通道；CUSTOM：自定义事件通道
	ProviderType *CreateChannelResponseProviderType `json:"provider_type,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateChannelResponse struct{}"
	}

	return strings.Join([]string{"CreateChannelResponse", string(data)}, " ")
}

type CreateChannelResponseProviderType struct {
	value string
}

type CreateChannelResponseProviderTypeEnum struct {
	OFFICIAL CreateChannelResponseProviderType
	CUSTOM   CreateChannelResponseProviderType
}

func GetCreateChannelResponseProviderTypeEnum() CreateChannelResponseProviderTypeEnum {
	return CreateChannelResponseProviderTypeEnum{
		OFFICIAL: CreateChannelResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: CreateChannelResponseProviderType{
			value: "CUSTOM",
		},
	}
}

func (c CreateChannelResponseProviderType) Value() string {
	return c.value
}

func (c CreateChannelResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateChannelResponseProviderType) UnmarshalJSON(b []byte) error {
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
