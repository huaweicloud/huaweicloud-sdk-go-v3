package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ConfigStatus **参数解释**: 采集通道配置状态 - OK 完成 - CHANGE 修改  **约束限制** 不涉及 **取值范围**: - OK - CHANGE  **默认值** 不涉及
type ConfigStatus struct {
	value string
}

type ConfigStatusEnum struct {
	OK     ConfigStatus
	CHANGE ConfigStatus
}

func GetConfigStatusEnum() ConfigStatusEnum {
	return ConfigStatusEnum{
		OK: ConfigStatus{
			value: "OK",
		},
		CHANGE: ConfigStatus{
			value: "CHANGE",
		},
	}
}

func (c ConfigStatus) Value() string {
	return c.value
}

func (c ConfigStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConfigStatus) UnmarshalJSON(b []byte) error {
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
