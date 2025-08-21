package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipelineStageJobDto 流水线阶段带任务
type PipelineStageJobDto struct {

	// 阶段ID
	Id *int32 `json:"id,omitempty"`

	// 阶段名称
	Name *string `json:"name,omitempty"`

	// 阶段顺序ID
	SortId *int32 `json:"sort_id,omitempty"`

	// 阶段状态, pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *PipelineStageJobDtoStatus `json:"status,omitempty"`

	Jobs *[]BaseJobDto `json:"jobs,omitempty"`
}

func (o PipelineStageJobDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineStageJobDto struct{}"
	}

	return strings.Join([]string{"PipelineStageJobDto", string(data)}, " ")
}

type PipelineStageJobDtoStatus struct {
	value string
}

type PipelineStageJobDtoStatusEnum struct {
	PENDING  PipelineStageJobDtoStatus
	RUNNING  PipelineStageJobDtoStatus
	SUCCESS  PipelineStageJobDtoStatus
	FAILED   PipelineStageJobDtoStatus
	CANCELED PipelineStageJobDtoStatus
	SKIPPED  PipelineStageJobDtoStatus
	TIMEDOUT PipelineStageJobDtoStatus
}

func GetPipelineStageJobDtoStatusEnum() PipelineStageJobDtoStatusEnum {
	return PipelineStageJobDtoStatusEnum{
		PENDING: PipelineStageJobDtoStatus{
			value: "pending",
		},
		RUNNING: PipelineStageJobDtoStatus{
			value: "running",
		},
		SUCCESS: PipelineStageJobDtoStatus{
			value: "success",
		},
		FAILED: PipelineStageJobDtoStatus{
			value: "failed",
		},
		CANCELED: PipelineStageJobDtoStatus{
			value: "canceled",
		},
		SKIPPED: PipelineStageJobDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: PipelineStageJobDtoStatus{
			value: "timedout",
		},
	}
}

func (c PipelineStageJobDtoStatus) Value() string {
	return c.value
}

func (c PipelineStageJobDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineStageJobDtoStatus) UnmarshalJSON(b []byte) error {
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
