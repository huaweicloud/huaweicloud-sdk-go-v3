package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// MainJob 实例复制等异步任务查询的响应体定义
type MainJob struct {

	// 任务状态
	Status *MainJobStatus `json:"status,omitempty"`

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

	// 子任务列表
	SubJobs *[]SubJob `json:"sub_jobs,omitempty"`
}

func (o MainJob) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MainJob struct{}"
	}

	return strings.Join([]string{"MainJob", string(data)}, " ")
}

type MainJobStatus struct {
	value string
}

type MainJobStatusEnum struct {
	INIT          MainJobStatus
	RUNNING       MainJobStatus
	FAIL          MainJobStatus
	SUCCESS       MainJobStatus
	ROLLBACKING   MainJobStatus
	COMPLETE      MainJobStatus
	ROLLBACK_FAIL MainJobStatus
	CANCEL        MainJobStatus
}

func GetMainJobStatusEnum() MainJobStatusEnum {
	return MainJobStatusEnum{
		INIT: MainJobStatus{
			value: "INIT",
		},
		RUNNING: MainJobStatus{
			value: "RUNNING",
		},
		FAIL: MainJobStatus{
			value: "FAIL",
		},
		SUCCESS: MainJobStatus{
			value: "SUCCESS",
		},
		ROLLBACKING: MainJobStatus{
			value: "ROLLBACKING",
		},
		COMPLETE: MainJobStatus{
			value: "COMPLETE",
		},
		ROLLBACK_FAIL: MainJobStatus{
			value: "ROLLBACK_FAIL",
		},
		CANCEL: MainJobStatus{
			value: "CANCEL",
		},
	}
}

func (c MainJobStatus) Value() string {
	return c.value
}

func (c MainJobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *MainJobStatus) UnmarshalJSON(b []byte) error {
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
