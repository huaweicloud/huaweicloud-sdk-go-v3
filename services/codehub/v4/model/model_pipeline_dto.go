package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipelineDto 流水线信息
type PipelineDto struct {

	// 流水线ID
	Id *int32 `json:"id,omitempty"`

	// 流水线链接
	WebUrl *string `json:"web_url,omitempty"`

	// commit id
	Sha *string `json:"sha,omitempty"`

	// 分支
	Ref *string `json:"ref,omitempty"`

	// 流水线状态，pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *PipelineDtoStatus `json:"status,omitempty"`

	// 流水线创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 流水线更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 流水线开始时间
	StartedAt *string `json:"started_at,omitempty"`

	// 流水线结束时间
	FinishedAt *string `json:"finished_at,omitempty"`
}

func (o PipelineDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineDto struct{}"
	}

	return strings.Join([]string{"PipelineDto", string(data)}, " ")
}

type PipelineDtoStatus struct {
	value string
}

type PipelineDtoStatusEnum struct {
	PENDING  PipelineDtoStatus
	RUNNING  PipelineDtoStatus
	SUCCESS  PipelineDtoStatus
	FAILED   PipelineDtoStatus
	CANCELED PipelineDtoStatus
	SKIPPED  PipelineDtoStatus
	TIMEDOUT PipelineDtoStatus
}

func GetPipelineDtoStatusEnum() PipelineDtoStatusEnum {
	return PipelineDtoStatusEnum{
		PENDING: PipelineDtoStatus{
			value: "pending",
		},
		RUNNING: PipelineDtoStatus{
			value: "running",
		},
		SUCCESS: PipelineDtoStatus{
			value: "success",
		},
		FAILED: PipelineDtoStatus{
			value: "failed",
		},
		CANCELED: PipelineDtoStatus{
			value: "canceled",
		},
		SKIPPED: PipelineDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: PipelineDtoStatus{
			value: "timedout",
		},
	}
}

func (c PipelineDtoStatus) Value() string {
	return c.value
}

func (c PipelineDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineDtoStatus) UnmarshalJSON(b []byte) error {
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
