package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SubscriptionSource struct {

	// 订阅源ID，需保证全局唯一。指定ID的订阅源存在时则进行更新，否则进行创建；未指定时由系统自动生成。由小写字母、数字、中划线组成，必须字母或数字开头。
	Id *string `json:"id,omitempty"`

	// 订阅的事件源名称
	Name string `json:"name"`

	// 订阅的事件源的提供方类型
	ProviderType SubscriptionSourceProviderType `json:"provider_type"`

	// 订阅的事件源参数列表, 该字段序列化后总长度不超过1024字节
	Detail *interface{} `json:"detail,omitempty"`

	// 订阅事件源的匹配过滤规则, 该字段序列化后总长度不超过2048字节
	Filter *interface{} `json:"filter"`
}

func (o SubscriptionSource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubscriptionSource struct{}"
	}

	return strings.Join([]string{"SubscriptionSource", string(data)}, " ")
}

type SubscriptionSourceProviderType struct {
	value string
}

type SubscriptionSourceProviderTypeEnum struct {
	CUSTOM   SubscriptionSourceProviderType
	OFFICIAL SubscriptionSourceProviderType
	PARTNER  SubscriptionSourceProviderType
}

func GetSubscriptionSourceProviderTypeEnum() SubscriptionSourceProviderTypeEnum {
	return SubscriptionSourceProviderTypeEnum{
		CUSTOM: SubscriptionSourceProviderType{
			value: "CUSTOM",
		},
		OFFICIAL: SubscriptionSourceProviderType{
			value: "OFFICIAL",
		},
		PARTNER: SubscriptionSourceProviderType{
			value: "PARTNER",
		},
	}
}

func (c SubscriptionSourceProviderType) Value() string {
	return c.value
}

func (c SubscriptionSourceProviderType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubscriptionSourceProviderType) UnmarshalJSON(b []byte) error {
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
