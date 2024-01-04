package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobType 训练类型。 * BASIC: 基础版(20句话) * MIDDLE: 进阶版(100句话) * ADVANCE: 高级版
type JobType struct {
	value string
}

type JobTypeEnum struct {
	BASIC   JobType
	MIDDLE  JobType
	ADVANCE JobType
}

func GetJobTypeEnum() JobTypeEnum {
	return JobTypeEnum{
		BASIC: JobType{
			value: "BASIC",
		},
		MIDDLE: JobType{
			value: "MIDDLE",
		},
		ADVANCE: JobType{
			value: "ADVANCE",
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
