package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobStatus **参数解释**: 作业状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
type JobStatus struct {
	value string
}

type JobStatusEnum struct {
	ENABLED  JobStatus
	DISABLED JobStatus
}

func GetJobStatusEnum() JobStatusEnum {
	return JobStatusEnum{
		ENABLED: JobStatus{
			value: "ENABLED",
		},
		DISABLED: JobStatus{
			value: "DISABLED",
		},
	}
}

func (c JobStatus) Value() string {
	return c.value
}

func (c JobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobStatus) UnmarshalJSON(b []byte) error {
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
