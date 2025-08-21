package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipelineStageDto 阶段信息
type PipelineStageDto struct {

	// 阶段ID
	Id *int32 `json:"id,omitempty"`

	// 仓库ID
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 流水线ID
	PipelineId *int32 `json:"pipeline_id,omitempty"`

	// 阶段名称
	Name *string `json:"name,omitempty"`

	// 阶段顺序id
	SortId *int32 `json:"sort_id,omitempty"`

	// 阶段状态, pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *PipelineStageDtoStatus `json:"status,omitempty"`
}

func (o PipelineStageDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineStageDto struct{}"
	}

	return strings.Join([]string{"PipelineStageDto", string(data)}, " ")
}

type PipelineStageDtoStatus struct {
	value string
}

type PipelineStageDtoStatusEnum struct {
	PENDING  PipelineStageDtoStatus
	RUNNING  PipelineStageDtoStatus
	SUCCESS  PipelineStageDtoStatus
	FAILED   PipelineStageDtoStatus
	CANCELED PipelineStageDtoStatus
	SKIPPED  PipelineStageDtoStatus
	TIMEDOUT PipelineStageDtoStatus
}

func GetPipelineStageDtoStatusEnum() PipelineStageDtoStatusEnum {
	return PipelineStageDtoStatusEnum{
		PENDING: PipelineStageDtoStatus{
			value: "pending",
		},
		RUNNING: PipelineStageDtoStatus{
			value: "running",
		},
		SUCCESS: PipelineStageDtoStatus{
			value: "success",
		},
		FAILED: PipelineStageDtoStatus{
			value: "failed",
		},
		CANCELED: PipelineStageDtoStatus{
			value: "canceled",
		},
		SKIPPED: PipelineStageDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: PipelineStageDtoStatus{
			value: "timedout",
		},
	}
}

func (c PipelineStageDtoStatus) Value() string {
	return c.value
}

func (c PipelineStageDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineStageDtoStatus) UnmarshalJSON(b []byte) error {
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
