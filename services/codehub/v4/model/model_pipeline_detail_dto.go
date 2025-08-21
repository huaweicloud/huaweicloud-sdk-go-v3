package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PipelineDetailDto 流水线详情
type PipelineDetailDto struct {

	// 流水线ID
	Id *int32 `json:"id,omitempty"`

	// 流水线链接
	WebUrl *string `json:"web_url,omitempty"`

	// commit id
	Sha *string `json:"sha,omitempty"`

	// 分支
	Ref *string `json:"ref,omitempty"`

	// 流水线状态，pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *PipelineDetailDtoStatus `json:"status,omitempty"`

	// 流水线创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 流水线更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 流水线开始时间
	StartedAt *string `json:"started_at,omitempty"`

	// 流水线结束时间
	FinishedAt *string `json:"finished_at,omitempty"`

	// 仓库ID
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 流水线是否失效
	IsInvalid *bool `json:"is_invalid,omitempty"`

	// 流水线类型，MERGE REQUEST代表为合并请求触发的流水线
	Type *PipelineDetailDtoType `json:"type,omitempty"`

	Stages *[]PipelineStageDto `json:"stages,omitempty"`

	// 是否是最近一条流水线
	IsLatest *bool `json:"is_latest,omitempty"`

	// 触发的用户
	TriggerUser *string `json:"trigger_user,omitempty"`

	// 是否所有job都运行完成
	AllJobFinished *bool `json:"all_job_finished,omitempty"`
}

func (o PipelineDetailDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineDetailDto struct{}"
	}

	return strings.Join([]string{"PipelineDetailDto", string(data)}, " ")
}

type PipelineDetailDtoStatus struct {
	value string
}

type PipelineDetailDtoStatusEnum struct {
	PENDING  PipelineDetailDtoStatus
	RUNNING  PipelineDetailDtoStatus
	SUCCESS  PipelineDetailDtoStatus
	FAILED   PipelineDetailDtoStatus
	CANCELED PipelineDetailDtoStatus
	SKIPPED  PipelineDetailDtoStatus
	TIMEDOUT PipelineDetailDtoStatus
}

func GetPipelineDetailDtoStatusEnum() PipelineDetailDtoStatusEnum {
	return PipelineDetailDtoStatusEnum{
		PENDING: PipelineDetailDtoStatus{
			value: "pending",
		},
		RUNNING: PipelineDetailDtoStatus{
			value: "running",
		},
		SUCCESS: PipelineDetailDtoStatus{
			value: "success",
		},
		FAILED: PipelineDetailDtoStatus{
			value: "failed",
		},
		CANCELED: PipelineDetailDtoStatus{
			value: "canceled",
		},
		SKIPPED: PipelineDetailDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: PipelineDetailDtoStatus{
			value: "timedout",
		},
	}
}

func (c PipelineDetailDtoStatus) Value() string {
	return c.value
}

func (c PipelineDetailDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineDetailDtoStatus) UnmarshalJSON(b []byte) error {
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

type PipelineDetailDtoType struct {
	value string
}

type PipelineDetailDtoTypeEnum struct {
	MERGE_REQUEST PipelineDetailDtoType
}

func GetPipelineDetailDtoTypeEnum() PipelineDetailDtoTypeEnum {
	return PipelineDetailDtoTypeEnum{
		MERGE_REQUEST: PipelineDetailDtoType{
			value: "MERGE REQUEST",
		},
	}
}

func (c PipelineDetailDtoType) Value() string {
	return c.value
}

func (c PipelineDetailDtoType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PipelineDetailDtoType) UnmarshalJSON(b []byte) error {
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
