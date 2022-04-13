package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 流水线资源信息
type PipelineCreationResult struct {
	Pipeline *PipelineBasic `json:"pipeline,omitempty"`
	// 任务id

	TaskId *string `json:"task_id,omitempty"`
	// 任务状态, success:成功,failed:失败,creating:创建中,cancel:取消,pending:等待创建

	Status *PipelineCreationResultStatus `json:"status,omitempty"`
	// 失败原因

	FailureReason *string `json:"failure_reason,omitempty"`
}

func (o PipelineCreationResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineCreationResult struct{}"
	}

	return strings.Join([]string{"PipelineCreationResult", string(data)}, " ")
}

type PipelineCreationResultStatus struct {
	value string
}

type PipelineCreationResultStatusEnum struct {
	SUCCESS  PipelineCreationResultStatus
	FAILED   PipelineCreationResultStatus
	CREATING PipelineCreationResultStatus
	CANCEL   PipelineCreationResultStatus
	PENDING  PipelineCreationResultStatus
}

func GetPipelineCreationResultStatusEnum() PipelineCreationResultStatusEnum {
	return PipelineCreationResultStatusEnum{
		SUCCESS: PipelineCreationResultStatus{
			value: "success",
		},
		FAILED: PipelineCreationResultStatus{
			value: "failed",
		},
		CREATING: PipelineCreationResultStatus{
			value: "creating",
		},
		CANCEL: PipelineCreationResultStatus{
			value: "cancel",
		},
		PENDING: PipelineCreationResultStatus{
			value: "pending",
		},
	}
}

func (c PipelineCreationResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineCreationResultStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
