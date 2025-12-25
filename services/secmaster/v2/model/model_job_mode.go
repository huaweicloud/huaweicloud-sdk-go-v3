package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobMode **参数解释**: 作业模式 - STREAMING 流式处理 - BATCH 批处理 - SEARCH 检索  **约束限制** 不涉及 **取值范围**: - STREAMING - BATCH - SEARCH  **默认值** 不涉及
type JobMode struct {
	value string
}

type JobModeEnum struct {
	STREAMING JobMode
	BATCH     JobMode
	SEARCH    JobMode
}

func GetJobModeEnum() JobModeEnum {
	return JobModeEnum{
		STREAMING: JobMode{
			value: "STREAMING",
		},
		BATCH: JobMode{
			value: "BATCH",
		},
		SEARCH: JobMode{
			value: "SEARCH",
		},
	}
}

func (c JobMode) Value() string {
	return c.value
}

func (c JobMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobMode) UnmarshalJSON(b []byte) error {
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
