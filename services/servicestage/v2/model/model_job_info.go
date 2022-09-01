package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 构建工程参数。
type JobInfo struct {

	// 创建者。
	CreatedBy *string `json:"CREATED_BY,omitempty" xml:"CREATED_BY"`

	// 执行状态。
	ExecutionStatus *JobInfoExecutionStatus `json:"EXECUTION_STATUS,omitempty" xml:"EXECUTION_STATUS"`

	// 工作描述。
	JobDesc *string `json:"JOB_DESC,omitempty" xml:"JOB_DESC"`

	// 工作ID。
	JobId *string `json:"JOB_ID,omitempty" xml:"JOB_ID"`

	// 工作名称。
	JobName *string `json:"JOB_NAME,omitempty" xml:"JOB_NAME"`

	// 类别。
	JobType *string `json:"JOB_TYPE,omitempty" xml:"JOB_TYPE"`

	// 排序ID。
	OrderId *string `json:"ORDER_ID,omitempty" xml:"ORDER_ID"`

	// 创建租户的项目ID。
	ProjectId *string `json:"PROJECT_ID,omitempty" xml:"PROJECT_ID"`

	// 实例ID。
	ServiceInstanceId *string `json:"SERVICE_INSTANCE_ID,omitempty" xml:"SERVICE_INSTANCE_ID"`
}

func (o JobInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInfo struct{}"
	}

	return strings.Join([]string{"JobInfo", string(data)}, " ")
}

type JobInfoExecutionStatus struct {
	value string
}

type JobInfoExecutionStatusEnum struct {
	RUNNING   JobInfoExecutionStatus
	FAILED    JobInfoExecutionStatus
	SUCCEEDED JobInfoExecutionStatus
}

func GetJobInfoExecutionStatusEnum() JobInfoExecutionStatusEnum {
	return JobInfoExecutionStatusEnum{
		RUNNING: JobInfoExecutionStatus{
			value: "RUNNING",
		},
		FAILED: JobInfoExecutionStatus{
			value: "FAILED",
		},
		SUCCEEDED: JobInfoExecutionStatus{
			value: "SUCCEEDED",
		},
	}
}

func (c JobInfoExecutionStatus) Value() string {
	return c.value
}

func (c JobInfoExecutionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInfoExecutionStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
