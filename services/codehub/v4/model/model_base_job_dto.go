package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseJobDto 任务详情
type BaseJobDto struct {

	// 任务ID
	Id *int32 `json:"id,omitempty"`

	// 提交ID
	Sha *string `json:"sha,omitempty"`

	// 分支
	Ref *string `json:"ref,omitempty"`

	// 阶段状态, pending为排队，running为运行中，success为成功，failed为失败，canceled为取消，skipped为跳过，timedout为超时
	Status *BaseJobDtoStatus `json:"status,omitempty"`

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
}

func (o BaseJobDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseJobDto struct{}"
	}

	return strings.Join([]string{"BaseJobDto", string(data)}, " ")
}

type BaseJobDtoStatus struct {
	value string
}

type BaseJobDtoStatusEnum struct {
	PENDING  BaseJobDtoStatus
	RUNNING  BaseJobDtoStatus
	SUCCESS  BaseJobDtoStatus
	FAILED   BaseJobDtoStatus
	CANCELED BaseJobDtoStatus
	SKIPPED  BaseJobDtoStatus
	TIMEDOUT BaseJobDtoStatus
}

func GetBaseJobDtoStatusEnum() BaseJobDtoStatusEnum {
	return BaseJobDtoStatusEnum{
		PENDING: BaseJobDtoStatus{
			value: "pending",
		},
		RUNNING: BaseJobDtoStatus{
			value: "running",
		},
		SUCCESS: BaseJobDtoStatus{
			value: "success",
		},
		FAILED: BaseJobDtoStatus{
			value: "failed",
		},
		CANCELED: BaseJobDtoStatus{
			value: "canceled",
		},
		SKIPPED: BaseJobDtoStatus{
			value: "skipped",
		},
		TIMEDOUT: BaseJobDtoStatus{
			value: "timedout",
		},
	}
}

func (c BaseJobDtoStatus) Value() string {
	return c.value
}

func (c BaseJobDtoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseJobDtoStatus) UnmarshalJSON(b []byte) error {
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
