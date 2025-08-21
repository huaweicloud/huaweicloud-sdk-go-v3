package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobWithStageDto 任务带阶段信息
type JobWithStageDto struct {

	// 任务ID
	Id *int32 `json:"id,omitempty"`

	// 提交ID
	Sha *string `json:"sha,omitempty"`

	// 分支
	Ref *string `json:"ref,omitempty"`

	// 阶段状态, pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *JobWithStageDtoStatus `json:"status,omitempty"`

	// 任务名称
	Name *string `json:"name,omitempty"`

	// 任务链接
	TargetUrl *string `json:"target_url,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 开始时间
	StartedAt *string `json:"started_at,omitempty"`

	// 结束时间
	FinishedAt *string `json:"finished_at,omitempty"`

	// 任务在构建系统中的ID
	ThirdBuildId *string `json:"third_build_id,omitempty"`

	// 阶段ID
	StageId *int32 `json:"stage_id,omitempty"`

	// 阶段名称
	Stage *string `json:"stage,omitempty"`
}

func (o JobWithStageDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobWithStageDto struct{}"
	}

	return strings.Join([]string{"JobWithStageDto", string(data)}, " ")
}

type JobWithStageDtoStatus struct {
	value string
}

type JobWithStageDtoStatusEnum struct {
	PENDING  JobWithStageDtoStatus
	RUNNING  JobWithStageDtoStatus
	SUCCESS  JobWithStageDtoStatus
	FAILED   JobWithStageDtoStatus
	CANCELED JobWithStageDtoStatus
	SKIPPED  JobWithStageDtoStatus
	TIMEDOUT JobWithStageDtoStatus
}

func GetJobWithStageDtoStatusEnum() JobWithStageDtoStatusEnum {
	return JobWithStageDtoStatusEnum{
		PENDING: JobWithStageDtoStatus{
			value: "pending",
		},
		RUNNING: JobWithStageDtoStatus{
			value: "running",
		},
		SUCCESS: JobWithStageDtoStatus{
			value: "success",
		},
		FAILED: JobWithStageDtoStatus{
			value: "failed",
		},
		CANCELED: JobWithStageDtoStatus{
			value: "canceled",
		},
		SKIPPED: JobWithStageDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: JobWithStageDtoStatus{
			value: "timedout",
		},
	}
}

func (c JobWithStageDtoStatus) Value() string {
	return c.value
}

func (c JobWithStageDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobWithStageDtoStatus) UnmarshalJSON(b []byte) error {
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
