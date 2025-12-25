package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlertRuleEnvironment **参数解释**: 环境类型 - PROD 生产环境 - TEST 测试环境  **约束限制** 不涉及 **取值范围**: - PROD - TEST  **默认值** 不涉及
type AlertRuleEnvironment struct {
	value string
}

type AlertRuleEnvironmentEnum struct {
	PROD AlertRuleEnvironment
	TEST AlertRuleEnvironment
}

func GetAlertRuleEnvironmentEnum() AlertRuleEnvironmentEnum {
	return AlertRuleEnvironmentEnum{
		PROD: AlertRuleEnvironment{
			value: "PROD",
		},
		TEST: AlertRuleEnvironment{
			value: "TEST",
		},
	}
}

func (c AlertRuleEnvironment) Value() string {
	return c.value
}

func (c AlertRuleEnvironment) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlertRuleEnvironment) UnmarshalJSON(b []byte) error {
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
