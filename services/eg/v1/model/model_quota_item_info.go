package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QuotaItemInfo struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 配额类型
	Type *QuotaItemInfoType `json:"type,omitempty"`

	// 配额最大值
	Max *int32 `json:"max,omitempty"`

	// 配额最小值
	Min *int32 `json:"min,omitempty"`

	// 当前租户的配额
	Quota *int32 `json:"quota,omitempty"`

	// 当前租户已使用的配额
	Used *int32 `json:"used,omitempty"`
}

func (o QuotaItemInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuotaItemInfo struct{}"
	}

	return strings.Join([]string{"QuotaItemInfo", string(data)}, " ")
}

type QuotaItemInfoType struct {
	value string
}

type QuotaItemInfoTypeEnum struct {
	CHANNEL              QuotaItemInfoType
	CHANNEL_SUBSCRIPTION QuotaItemInfoType
	SOURCE               QuotaItemInfoType
	SUBSCRIPTION         QuotaItemInfoType
	SUBSCRIPTION_TARGET  QuotaItemInfoType
	SCHEMA               QuotaItemInfoType
	SCHEMA_VERSION       QuotaItemInfoType
	CONNECTION           QuotaItemInfoType
	PRIVATE_ENDPOINT     QuotaItemInfoType
	SOURCE_RABBITMQ      QuotaItemInfoType
	SOURCE_ROCKETMQ      QuotaItemInfoType
}

func GetQuotaItemInfoTypeEnum() QuotaItemInfoTypeEnum {
	return QuotaItemInfoTypeEnum{
		CHANNEL: QuotaItemInfoType{
			value: "CHANNEL",
		},
		CHANNEL_SUBSCRIPTION: QuotaItemInfoType{
			value: "CHANNEL_SUBSCRIPTION",
		},
		SOURCE: QuotaItemInfoType{
			value: "SOURCE",
		},
		SUBSCRIPTION: QuotaItemInfoType{
			value: "SUBSCRIPTION",
		},
		SUBSCRIPTION_TARGET: QuotaItemInfoType{
			value: "SUBSCRIPTION_TARGET",
		},
		SCHEMA: QuotaItemInfoType{
			value: "SCHEMA",
		},
		SCHEMA_VERSION: QuotaItemInfoType{
			value: "SCHEMA_VERSION",
		},
		CONNECTION: QuotaItemInfoType{
			value: "CONNECTION",
		},
		PRIVATE_ENDPOINT: QuotaItemInfoType{
			value: "PRIVATE_ENDPOINT",
		},
		SOURCE_RABBITMQ: QuotaItemInfoType{
			value: "SOURCE_RABBITMQ",
		},
		SOURCE_ROCKETMQ: QuotaItemInfoType{
			value: "SOURCE_ROCKETMQ",
		},
	}
}

func (c QuotaItemInfoType) Value() string {
	return c.value
}

func (c QuotaItemInfoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *QuotaItemInfoType) UnmarshalJSON(b []byte) error {
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
