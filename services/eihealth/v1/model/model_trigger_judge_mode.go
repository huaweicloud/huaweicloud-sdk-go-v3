package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TriggerJudgeMode 触发器的判断模式
type TriggerJudgeMode struct {
	value string
}

type TriggerJudgeModeEnum struct {
	GTE      TriggerJudgeMode
	LTE      TriggerJudgeMode
	GT       TriggerJudgeMode
	LT       TriggerJudgeMode
	LIKE     TriggerJudgeMode
	NOTLIKE  TriggerJudgeMode
	EQUAL    TriggerJudgeMode
	NOTEQUAL TriggerJudgeMode
}

func GetTriggerJudgeModeEnum() TriggerJudgeModeEnum {
	return TriggerJudgeModeEnum{
		GTE: TriggerJudgeMode{
			value: "gte",
		},
		LTE: TriggerJudgeMode{
			value: "lte",
		},
		GT: TriggerJudgeMode{
			value: "gt",
		},
		LT: TriggerJudgeMode{
			value: "lt",
		},
		LIKE: TriggerJudgeMode{
			value: "like",
		},
		NOTLIKE: TriggerJudgeMode{
			value: "notlike",
		},
		EQUAL: TriggerJudgeMode{
			value: "equal",
		},
		NOTEQUAL: TriggerJudgeMode{
			value: "notequal",
		},
	}
}

func (c TriggerJudgeMode) Value() string {
	return c.value
}

func (c TriggerJudgeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerJudgeMode) UnmarshalJSON(b []byte) error {
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
