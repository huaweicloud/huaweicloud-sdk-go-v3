package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigurationPrimitiveTypeHolderConfiguration 私有hook配置项，可以指定私有hook生效的目标资源栈以及私有hook校验失败后资源栈的行为。可通过UpdatePrivateHook API更新私有hook配置项。
type ConfigurationPrimitiveTypeHolderConfiguration struct {

	// 指定私有hook生效的目标资源栈，有效值为NONE或ALL。  NONE：指定此私有hook不会作用于任何资源栈 ALL：指定此私有hook将会应用于账号下的所有资源栈
	TargetStacks *ConfigurationPrimitiveTypeHolderConfigurationTargetStacks `json:"target_stacks,omitempty"`

	// 指定私有hook校验失败后的行为，有效值为FAIL或WARN。  FAIL：指定此私有hook校验失败后资源栈将停止部署，资源栈状态将更新为DEPLOYMENT_FAILED。 WARN：指定此私有hook校验失败后仅通过资源栈事件展示警告消息，但不影响资源栈部署。
	FailureMode *ConfigurationPrimitiveTypeHolderConfigurationFailureMode `json:"failure_mode,omitempty"`
}

func (o ConfigurationPrimitiveTypeHolderConfiguration) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigurationPrimitiveTypeHolderConfiguration struct{}"
	}

	return strings.Join([]string{"ConfigurationPrimitiveTypeHolderConfiguration", string(data)}, " ")
}

type ConfigurationPrimitiveTypeHolderConfigurationTargetStacks struct {
	value string
}

type ConfigurationPrimitiveTypeHolderConfigurationTargetStacksEnum struct {
	NONE ConfigurationPrimitiveTypeHolderConfigurationTargetStacks
	ALL  ConfigurationPrimitiveTypeHolderConfigurationTargetStacks
}

func GetConfigurationPrimitiveTypeHolderConfigurationTargetStacksEnum() ConfigurationPrimitiveTypeHolderConfigurationTargetStacksEnum {
	return ConfigurationPrimitiveTypeHolderConfigurationTargetStacksEnum{
		NONE: ConfigurationPrimitiveTypeHolderConfigurationTargetStacks{
			value: "NONE",
		},
		ALL: ConfigurationPrimitiveTypeHolderConfigurationTargetStacks{
			value: "ALL",
		},
	}
}

func (c ConfigurationPrimitiveTypeHolderConfigurationTargetStacks) Value() string {
	return c.value
}

func (c ConfigurationPrimitiveTypeHolderConfigurationTargetStacks) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationPrimitiveTypeHolderConfigurationTargetStacks) UnmarshalJSON(b []byte) error {
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

type ConfigurationPrimitiveTypeHolderConfigurationFailureMode struct {
	value string
}

type ConfigurationPrimitiveTypeHolderConfigurationFailureModeEnum struct {
	WARN ConfigurationPrimitiveTypeHolderConfigurationFailureMode
	FAIL ConfigurationPrimitiveTypeHolderConfigurationFailureMode
}

func GetConfigurationPrimitiveTypeHolderConfigurationFailureModeEnum() ConfigurationPrimitiveTypeHolderConfigurationFailureModeEnum {
	return ConfigurationPrimitiveTypeHolderConfigurationFailureModeEnum{
		WARN: ConfigurationPrimitiveTypeHolderConfigurationFailureMode{
			value: "WARN",
		},
		FAIL: ConfigurationPrimitiveTypeHolderConfigurationFailureMode{
			value: "FAIL",
		},
	}
}

func (c ConfigurationPrimitiveTypeHolderConfigurationFailureMode) Value() string {
	return c.value
}

func (c ConfigurationPrimitiveTypeHolderConfigurationFailureMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigurationPrimitiveTypeHolderConfigurationFailureMode) UnmarshalJSON(b []byte) error {
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
