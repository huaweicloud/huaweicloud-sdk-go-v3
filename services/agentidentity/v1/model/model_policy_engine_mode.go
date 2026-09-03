package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PolicyEngineMode 策略集的强制模式。LOG_ONLY - 策略集会根据您的策略评估每个操作并添加工具调用是否允许或拒绝的跟踪，但不强制执行决策。使用此模式在启用强制执行之前测试和验证策略。ENFORCE - 策略集会根据您的策略评估操作并通过允许或拒绝代理操作来强制执行决策。在启用强制执行之前，请在 LOG_ONLY 模式下测试和验证策略，以避免意外拒绝或对生产流量产生不利影响。
type PolicyEngineMode struct {
	value string
}

type PolicyEngineModeEnum struct {
	LOG_ONLY PolicyEngineMode
	ENFORCE  PolicyEngineMode
}

func GetPolicyEngineModeEnum() PolicyEngineModeEnum {
	return PolicyEngineModeEnum{
		LOG_ONLY: PolicyEngineMode{
			value: "LOG_ONLY",
		},
		ENFORCE: PolicyEngineMode{
			value: "ENFORCE",
		},
	}
}

func (c PolicyEngineMode) Value() string {
	return c.value
}

func (c PolicyEngineMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PolicyEngineMode) UnmarshalJSON(b []byte) error {
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
