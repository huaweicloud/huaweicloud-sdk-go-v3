package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ToolType **参数解释**： 实验类型。 **约束限制**： 不涉及 **取值范围**： * drug_job：药物实验作业 * workflow_job：流程实验作业 **默认取值**： 不涉及
type ToolType struct {
	value string
}

type ToolTypeEnum struct {
	DRUG_JOB     ToolType
	WORKFLOW_JOB ToolType
}

func GetToolTypeEnum() ToolTypeEnum {
	return ToolTypeEnum{
		DRUG_JOB: ToolType{
			value: "drug_job",
		},
		WORKFLOW_JOB: ToolType{
			value: "workflow_job",
		},
	}
}

func (c ToolType) Value() string {
	return c.value
}

func (c ToolType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ToolType) UnmarshalJSON(b []byte) error {
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
