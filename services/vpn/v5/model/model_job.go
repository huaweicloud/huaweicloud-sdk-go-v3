package model

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"strings"
)

// Job 任务
type Job struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 资源Id
	ResourceId *string `json:"resource_id,omitempty"`

	// 任务类型
	JobType *JobJobType `json:"job_type,omitempty"`

	// 任务状态
	Status *JobStatus `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 完成时间
	FinishedAt *sdktime.SdkTime `json:"finished_at,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 子任务列表
	SubJobs *[]SubJob `json:"sub_jobs,omitempty"`
}

func (o Job) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Job struct{}"
	}

	return strings.Join([]string{"Job", string(data)}, " ")
}

type JobJobType struct {
	value string
}

type JobJobTypeEnum struct {
	UPGRADE  JobJobType
	ROLLBACK JobJobType
}

func GetJobJobTypeEnum() JobJobTypeEnum {
	return JobJobTypeEnum{
		UPGRADE: JobJobType{
			value: "upgrade",
		},
		ROLLBACK: JobJobType{
			value: "rollback",
		},
	}
}

func (c JobJobType) Value() string {
	return c.value
}

func (c JobJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobJobType) UnmarshalJSON(b []byte) error {
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

type JobStatus struct {
	value string
}

type JobStatusEnum struct {
	UPGRADING               JobStatus
	PENDING_UPGRADE_CONFIRM JobStatus
	SUCCESS                 JobStatus
	ROLLING_BACK            JobStatus
	ROLLBACK_SUCCESS        JobStatus
	FAIL                    JobStatus
}

func GetJobStatusEnum() JobStatusEnum {
	return JobStatusEnum{
		UPGRADING: JobStatus{
			value: "upgrading",
		},
		PENDING_UPGRADE_CONFIRM: JobStatus{
			value: "pending_upgrade_confirm",
		},
		SUCCESS: JobStatus{
			value: "success",
		},
		ROLLING_BACK: JobStatus{
			value: "rolling_back",
		},
		ROLLBACK_SUCCESS: JobStatus{
			value: "rollback_success",
		},
		FAIL: JobStatus{
			value: "fail",
		},
	}
}

func (c JobStatus) Value() string {
	return c.value
}

func (c JobStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobStatus) UnmarshalJSON(b []byte) error {
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
