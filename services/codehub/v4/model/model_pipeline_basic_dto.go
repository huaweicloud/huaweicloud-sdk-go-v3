package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipelineBasicDto 流水线基本信息
type PipelineBasicDto struct {

	// 流水线ID
	Id *int32 `json:"id,omitempty"`

	// 流水线链接
	WebUrl *string `json:"web_url,omitempty"`

	// commit id
	Sha *string `json:"sha,omitempty"`

	// 分支
	Ref *string `json:"ref,omitempty"`

	// 流水线状态，pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *PipelineBasicDtoStatus `json:"status,omitempty"`
}

func (o PipelineBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineBasicDto struct{}"
	}

	return strings.Join([]string{"PipelineBasicDto", string(data)}, " ")
}

type PipelineBasicDtoStatus struct {
	value string
}

type PipelineBasicDtoStatusEnum struct {
	PENDING  PipelineBasicDtoStatus
	RUNNING  PipelineBasicDtoStatus
	SUCCESS  PipelineBasicDtoStatus
	FAILED   PipelineBasicDtoStatus
	CANCELED PipelineBasicDtoStatus
	SKIPPED  PipelineBasicDtoStatus
	TIMEDOUT PipelineBasicDtoStatus
}

func GetPipelineBasicDtoStatusEnum() PipelineBasicDtoStatusEnum {
	return PipelineBasicDtoStatusEnum{
		PENDING: PipelineBasicDtoStatus{
			value: "pending",
		},
		RUNNING: PipelineBasicDtoStatus{
			value: "running",
		},
		SUCCESS: PipelineBasicDtoStatus{
			value: "success",
		},
		FAILED: PipelineBasicDtoStatus{
			value: "failed",
		},
		CANCELED: PipelineBasicDtoStatus{
			value: "canceled",
		},
		SKIPPED: PipelineBasicDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: PipelineBasicDtoStatus{
			value: "timedout",
		},
	}
}

func (c PipelineBasicDtoStatus) Value() string {
	return c.value
}

func (c PipelineBasicDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineBasicDtoStatus) UnmarshalJSON(b []byte) error {
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
