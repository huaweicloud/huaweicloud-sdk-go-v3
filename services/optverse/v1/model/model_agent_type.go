package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgentType **参数解释**： 助手类型。 **约束限制**： 不涉及 **取值范围**： * optverse：天筹工具链 * predict：预测工具链 * vrp：路径规划工具链 * demand：需求助手 **默认取值**： 不涉及
type AgentType struct {
	value string
}

type AgentTypeEnum struct {
	OPTVERSE AgentType
	PREDICT  AgentType
	VRP      AgentType
	DEMAND   AgentType
}

func GetAgentTypeEnum() AgentTypeEnum {
	return AgentTypeEnum{
		OPTVERSE: AgentType{
			value: "optverse",
		},
		PREDICT: AgentType{
			value: "predict",
		},
		VRP: AgentType{
			value: "vrp",
		},
		DEMAND: AgentType{
			value: "demand",
		},
	}
}

func (c AgentType) Value() string {
	return c.value
}

func (c AgentType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AgentType) UnmarshalJSON(b []byte) error {
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
