package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobType 训练类型。 * LLM_MAJOR: 大模型中英文版 * LLM_MINOR: 大模型中小语种版 * BASIC: 基础版(20句话) * MIDDLE: 进阶版(100句话) * ADVANCE: 高级版 * THIRD_PARTY: 第三方训练版
type JobType struct {
	value string
}

type JobTypeEnum struct {
	LLM_MAJOR   JobType
	LLM_MINOR   JobType
	BASIC       JobType
	MIDDLE      JobType
	ADVANCE     JobType
	THIRD_PARTY JobType
}

func GetJobTypeEnum() JobTypeEnum {
	return JobTypeEnum{
		LLM_MAJOR: JobType{
			value: "LLM_MAJOR",
		},
		LLM_MINOR: JobType{
			value: "LLM_MINOR",
		},
		BASIC: JobType{
			value: "BASIC",
		},
		MIDDLE: JobType{
			value: "MIDDLE",
		},
		ADVANCE: JobType{
			value: "ADVANCE",
		},
		THIRD_PARTY: JobType{
			value: "THIRD_PARTY",
		},
	}
}

func (c JobType) Value() string {
	return c.value
}

func (c JobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobType) UnmarshalJSON(b []byte) error {
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
