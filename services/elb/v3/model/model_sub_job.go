package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// SubJob 子任务
type SubJob struct {

	// 任务状态
	Status *SubJobStatus `json:"status,omitempty"`

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

	// 参数解释：资源ID。
	ResourceId *string `json:"resource_id,omitempty"`

	// 任务关联的资源列表
	Entities *[]JobEntities `json:"entities,omitempty"`
}

func (o SubJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubJob struct{}"
	}

	return strings.Join([]string{"SubJob", string(data)}, " ")
}

type SubJobStatus struct {
	value string
}

type SubJobStatusEnum struct {
	INIT          SubJobStatus
	RUNNING       SubJobStatus
	FAIL          SubJobStatus
	SUCCESS       SubJobStatus
	ROLLBACKING   SubJobStatus
	COMPLETE      SubJobStatus
	ROLLBACK_FAIL SubJobStatus
	CANCEL        SubJobStatus
}

func GetSubJobStatusEnum() SubJobStatusEnum {
	return SubJobStatusEnum{
		INIT: SubJobStatus{
			value: "INIT",
		},
		RUNNING: SubJobStatus{
			value: "RUNNING",
		},
		FAIL: SubJobStatus{
			value: "FAIL",
		},
		SUCCESS: SubJobStatus{
			value: "SUCCESS",
		},
		ROLLBACKING: SubJobStatus{
			value: "ROLLBACKING",
		},
		COMPLETE: SubJobStatus{
			value: "COMPLETE",
		},
		ROLLBACK_FAIL: SubJobStatus{
			value: "ROLLBACK_FAIL",
		},
		CANCEL: SubJobStatus{
			value: "CANCEL",
		},
	}
}

func (c SubJobStatus) Value() string {
	return c.value
}

func (c SubJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SubJobStatus) UnmarshalJSON(b []byte) error {
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
