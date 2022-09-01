package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type QuotaItemInfo struct {

	// 名称
	Name *string `json:"name,omitempty" xml:"name"`

	// 配额类型
	Type *QuotaItemInfoType `json:"type,omitempty" xml:"type"`

	// 配额最大值
	Max *string `json:"max,omitempty" xml:"max"`

	// 配额最小值
	Min *string `json:"min,omitempty" xml:"min"`

	// 当前租户的配额
	Quota *string `json:"quota,omitempty" xml:"quota"`

	// 当前租户已使用的配额
	Used *string `json:"used,omitempty" xml:"used"`
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
