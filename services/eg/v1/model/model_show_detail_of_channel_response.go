package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowDetailOfChannelResponse Response Object
type ShowDetailOfChannelResponse struct {

	// 通道ID
	Id *string `json:"id,omitempty"`

	// 通道名称
	Name *string `json:"name,omitempty"`

	// 通道描述
	Description *string `json:"description,omitempty"`

	// 通道提供方类型，OFFICIAL：官方事件通道；CUSTOM：自定义事件通道；PARTNER：伙伴事件通道
	ProviderType *ShowDetailOfChannelResponseProviderType `json:"provider_type,omitempty"`

	// 创建UTC时间
	CreatedTime *string `json:"created_time,omitempty"`

	// 更新UTC时间
	UpdatedTime    *string `json:"updated_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDetailOfChannelResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailOfChannelResponse struct{}"
	}

	return strings.Join([]string{"ShowDetailOfChannelResponse", string(data)}, " ")
}

type ShowDetailOfChannelResponseProviderType struct {
	value string
}

type ShowDetailOfChannelResponseProviderTypeEnum struct {
	OFFICIAL ShowDetailOfChannelResponseProviderType
	CUSTOM   ShowDetailOfChannelResponseProviderType
	PARTNER  ShowDetailOfChannelResponseProviderType
}

func GetShowDetailOfChannelResponseProviderTypeEnum() ShowDetailOfChannelResponseProviderTypeEnum {
	return ShowDetailOfChannelResponseProviderTypeEnum{
		OFFICIAL: ShowDetailOfChannelResponseProviderType{
			value: "OFFICIAL",
		},
		CUSTOM: ShowDetailOfChannelResponseProviderType{
			value: "CUSTOM",
		},
		PARTNER: ShowDetailOfChannelResponseProviderType{
			value: "PARTNER",
		},
	}
}

func (c ShowDetailOfChannelResponseProviderType) Value() string {
	return c.value
}

func (c ShowDetailOfChannelResponseProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowDetailOfChannelResponseProviderType) UnmarshalJSON(b []byte) error {
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
