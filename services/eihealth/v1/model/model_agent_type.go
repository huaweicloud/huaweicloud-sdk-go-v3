package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AgentType **参数解释**： 助手类型。 **约束限制**： 不涉及 **取值范围**： * auto：通用助手 * drug：药物助手 * gene：基因助手 **默认取值**： 不涉及
type AgentType struct {
	value string
}

type AgentTypeEnum struct {
	AUTO AgentType
	GENE AgentType
	DRUG AgentType
}

func GetAgentTypeEnum() AgentTypeEnum {
	return AgentTypeEnum{
		AUTO: AgentType{
			value: "auto",
		},
		GENE: AgentType{
			value: "gene",
		},
		DRUG: AgentType{
			value: "drug",
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
