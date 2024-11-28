package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StandardPlatformApiConfig 自定义直播平台回调配置
type StandardPlatformApiConfig struct {

	// API类型
	ApiType StandardPlatformApiConfigApiType `json:"api_type"`

	// URL。仅支持HTTPS形式URL
	Url string `json:"url"`
}

func (o StandardPlatformApiConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StandardPlatformApiConfig struct{}"
	}

	return strings.Join([]string{"StandardPlatformApiConfig", string(data)}, " ")
}

type StandardPlatformApiConfigApiType struct {
	value string
}

type StandardPlatformApiConfigApiTypeEnum struct {
	PRODUCT_QUERY       StandardPlatformApiConfigApiType
	LIVE_EVENT_CALLBACK StandardPlatformApiConfigApiType
}

func GetStandardPlatformApiConfigApiTypeEnum() StandardPlatformApiConfigApiTypeEnum {
	return StandardPlatformApiConfigApiTypeEnum{
		PRODUCT_QUERY: StandardPlatformApiConfigApiType{
			value: "PRODUCT_QUERY",
		},
		LIVE_EVENT_CALLBACK: StandardPlatformApiConfigApiType{
			value: "LIVE_EVENT_CALLBACK",
		},
	}
}

func (c StandardPlatformApiConfigApiType) Value() string {
	return c.value
}

func (c StandardPlatformApiConfigApiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StandardPlatformApiConfigApiType) UnmarshalJSON(b []byte) error {
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
