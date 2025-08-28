package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PluginProviderEnum 插件供应商 * AMAP_WEATHER：高德天气 * BOCHA：博查
type PluginProviderEnum struct {
	value string
}

type PluginProviderEnumEnum struct {
	AMAP_WEATHER PluginProviderEnum
	BOCHA        PluginProviderEnum
}

func GetPluginProviderEnumEnum() PluginProviderEnumEnum {
	return PluginProviderEnumEnum{
		AMAP_WEATHER: PluginProviderEnum{
			value: "AMAP_WEATHER",
		},
		BOCHA: PluginProviderEnum{
			value: "BOCHA",
		},
	}
}

func (c PluginProviderEnum) Value() string {
	return c.value
}

func (c PluginProviderEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginProviderEnum) UnmarshalJSON(b []byte) error {
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
