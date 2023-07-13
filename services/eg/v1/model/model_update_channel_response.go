package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateChannelResponse Response Object
type UpdateChannelResponse struct {

	// 通道ID
	Id *string `json:"id,omitempty"`

	// 通道名称
	Name *string `json:"name,omitempty"`

	// 通道描述
	Description *string `json:"description,omitempty"`

	// 通道提供方类型，OFFICIAL：官方事件通道；CUSTOM：自定义事件通道；PARTNER：伙伴事件通道
	ProviderType *UpdateChannelResponseProviderType `json:"provider_type,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime *string `json:"updated_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateChannelResponse struct{}"
	}

	return strings.Join([]string{"UpdateChannelResponse", string(data)}, " ")
}

type UpdateChannelResponseProviderType struct {
	value string
}

type UpdateChannelResponseProviderTypeEnum struct {
	OFFICIAL UpdateChannelResponseProviderType
	CUSTOM   UpdateChannelResponseProviderType
	PARTNER  UpdateChannelResponseProviderType
}

func GetUpdateChannelResponseProviderTypeEnum() UpdateChannelResponseProviderTypeEnum {
	return UpdateChannelResponseProviderTypeEnum{
		OFFICIAL: UpdateChannelResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: UpdateChannelResponseProviderType{
			value: "CUSTOM",
		},
		PARTNER: UpdateChannelResponseProviderType{
			value: "PARTNER",
		},
	}
}

func (c UpdateChannelResponseProviderType) Value() string {
	return c.value
}

func (c UpdateChannelResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateChannelResponseProviderType) UnmarshalJSON(b []byte) error {
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
