package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgentRole **参数解释**： 助手角色。 **约束限制**： 不涉及 **取值范围**： * Common：对话助手角色 * Biomed：作业助手角色 **默认取值**： Common
type AgentRole struct {
	value string
}

type AgentRoleEnum struct {
	COMMON AgentRole
	BIOMED AgentRole
}

func GetAgentRoleEnum() AgentRoleEnum {
	return AgentRoleEnum{
		COMMON: AgentRole{
			value: "Common",
		},
		BIOMED: AgentRole{
			value: "Biomed",
		},
	}
}

func (c AgentRole) Value() string {
	return c.value
}

func (c AgentRole) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentRole) UnmarshalJSON(b []byte) error {
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
