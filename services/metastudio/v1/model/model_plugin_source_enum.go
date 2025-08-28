package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PluginSourceEnum 使用的插件来源 * PLUGIN_CONFIG：插件配置 * DEFAULT：默认
type PluginSourceEnum struct {
	value string
}

type PluginSourceEnumEnum struct {
	PLUGIN_CONFIG PluginSourceEnum
	DEFAULT       PluginSourceEnum
}

func GetPluginSourceEnumEnum() PluginSourceEnumEnum {
	return PluginSourceEnumEnum{
		PLUGIN_CONFIG: PluginSourceEnum{
			value: "PLUGIN_CONFIG",
		},
		DEFAULT: PluginSourceEnum{
			value: "DEFAULT",
		},
	}
}

func (c PluginSourceEnum) Value() string {
	return c.value
}

func (c PluginSourceEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PluginSourceEnum) UnmarshalJSON(b []byte) error {
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
