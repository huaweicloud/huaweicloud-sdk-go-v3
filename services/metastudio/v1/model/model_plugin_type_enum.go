package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PluginTypeEnum 插件类型 * WEATHER_QUERY：天气查询 * WEB_SEARCH：网络搜索
type PluginTypeEnum struct {
	value string
}

type PluginTypeEnumEnum struct {
	WEATHER_QUERY PluginTypeEnum
	WEB_SEARCH    PluginTypeEnum
}

func GetPluginTypeEnumEnum() PluginTypeEnumEnum {
	return PluginTypeEnumEnum{
		WEATHER_QUERY: PluginTypeEnum{
			value: "WEATHER_QUERY",
		},
		WEB_SEARCH: PluginTypeEnum{
			value: "WEB_SEARCH",
		},
	}
}

func (c PluginTypeEnum) Value() string {
	return c.value
}

func (c PluginTypeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginTypeEnum) UnmarshalJSON(b []byte) error {
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
