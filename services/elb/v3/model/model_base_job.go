package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BaseJob 任务的基类
type BaseJob struct {

	// 任务状态
	Status *BaseJobStatus `json:"status,omitempty"`

	// 任务开始时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 任务结束时间
	EndTime *string `json:"end_time,omitempty"`

	// 任务ID
	JobId *string `json:"job_id,omitempty"`

	// 任务类型
	JobType *string `json:"job_type,omitempty"`

	// 任务的错误码
	ErrorCode *string `json:"error_code,omitempty"`

	// 任务的错误信息
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：资源ID。
	ResourceId *string `json:"resource_id,omitempty"`
}

func (o BaseJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseJob struct{}"
	}

	return strings.Join([]string{"BaseJob", string(data)}, " ")
}

type BaseJobStatus struct {
	value string
}

type BaseJobStatusEnum struct {
	INIT          BaseJobStatus
	RUNNING       BaseJobStatus
	FAIL          BaseJobStatus
	SUCCESS       BaseJobStatus
	ROLLBACKING   BaseJobStatus
	COMPLETE      BaseJobStatus
	ROLLBACK_FAIL BaseJobStatus
	CANCEL        BaseJobStatus
}

func GetBaseJobStatusEnum() BaseJobStatusEnum {
	return BaseJobStatusEnum{
		INIT: BaseJobStatus{
			value: "INIT",
		},
		RUNNING: BaseJobStatus{
			value: "RUNNING",
		},
		FAIL: BaseJobStatus{
			value: "FAIL",
		},
		SUCCESS: BaseJobStatus{
			value: "SUCCESS",
		},
		ROLLBACKING: BaseJobStatus{
			value: "ROLLBACKING",
		},
		COMPLETE: BaseJobStatus{
			value: "COMPLETE",
		},
		ROLLBACK_FAIL: BaseJobStatus{
			value: "ROLLBACK_FAIL",
		},
		CANCEL: BaseJobStatus{
			value: "CANCEL",
		},
	}
}

func (c BaseJobStatus) Value() string {
	return c.value
}

func (c BaseJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BaseJobStatus) UnmarshalJSON(b []byte) error {
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
