package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobProgress 实例初始化进度。
type JobProgress struct {

	// **参数解释**：任务某个步骤的状态。 **取值范围**：枚举类型，取值如下： - WAITING：等待中 - PROCESSING：处理中 - FAILED：任务失败 - COMPLETED：任务完成
	Status *JobProgressStatus `json:"status,omitempty"`

	// **参数解释**：任务的步骤。 **取值范围**：枚举类型，取值如下： - 1：准备存储 - 2：准备计算资源 - 3：配置网络 - 4：初始化实例
	Step *int32 `json:"step,omitempty"`

	// **参数解释**：任务某个步骤的描述。 **取值范围**：不涉及。
	Description *string `json:"description,omitempty"`
}

func (o JobProgress) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobProgress struct{}"
	}

	return strings.Join([]string{"JobProgress", string(data)}, " ")
}

type JobProgressStatus struct {
	value string
}

type JobProgressStatusEnum struct {
	COMPLETED  JobProgressStatus
	FAILED     JobProgressStatus
	PROCESSING JobProgressStatus
	WAITING    JobProgressStatus
}

func GetJobProgressStatusEnum() JobProgressStatusEnum {
	return JobProgressStatusEnum{
		COMPLETED: JobProgressStatus{
			value: "COMPLETED",
		},
		FAILED: JobProgressStatus{
			value: "FAILED",
		},
		PROCESSING: JobProgressStatus{
			value: "PROCESSING",
		},
		WAITING: JobProgressStatus{
			value: "WAITING",
		},
	}
}

func (c JobProgressStatus) Value() string {
	return c.value
}

func (c JobProgressStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobProgressStatus) UnmarshalJSON(b []byte) error {
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
