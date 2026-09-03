package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstallStatusEnum 安装状态枚举。
type InstallStatusEnum struct {
	value string
}

type InstallStatusEnumEnum struct {
	INSTALLING   InstallStatusEnum
	INSTALLED    InstallStatusEnum
	FAILED       InstallStatusEnum
	UNINSTALLING InstallStatusEnum
	UNINSTALLED  InstallStatusEnum
}

func GetInstallStatusEnumEnum() InstallStatusEnumEnum {
	return InstallStatusEnumEnum{
		INSTALLING: InstallStatusEnum{
			value: "INSTALLING",
		},
		INSTALLED: InstallStatusEnum{
			value: "INSTALLED",
		},
		FAILED: InstallStatusEnum{
			value: "FAILED",
		},
		UNINSTALLING: InstallStatusEnum{
			value: "UNINSTALLING",
		},
		UNINSTALLED: InstallStatusEnum{
			value: "UNINSTALLED",
		},
	}
}

func (c InstallStatusEnum) Value() string {
	return c.value
}

func (c InstallStatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstallStatusEnum) UnmarshalJSON(b []byte) error {
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
