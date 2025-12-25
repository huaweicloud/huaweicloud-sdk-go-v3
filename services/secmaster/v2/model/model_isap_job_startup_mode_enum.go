package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// IsapJobStartupModeEnum **参数解释**: 作业启动模式 - UPGRADE 升级启动 - REFRESH_NEW 全新启动  **约束限制** 不涉及 **取值范围**: - UPGRADE - REFRESH_NEW  **默认值** 不涉及
type IsapJobStartupModeEnum struct {
	value string
}

type IsapJobStartupModeEnumEnum struct {
	UPGRADE     IsapJobStartupModeEnum
	REFRESH_NEW IsapJobStartupModeEnum
}

func GetIsapJobStartupModeEnumEnum() IsapJobStartupModeEnumEnum {
	return IsapJobStartupModeEnumEnum{
		UPGRADE: IsapJobStartupModeEnum{
			value: "UPGRADE",
		},
		REFRESH_NEW: IsapJobStartupModeEnum{
			value: "REFRESH_NEW",
		},
	}
}

func (c IsapJobStartupModeEnum) Value() string {
	return c.value
}

func (c IsapJobStartupModeEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *IsapJobStartupModeEnum) UnmarshalJSON(b []byte) error {
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
