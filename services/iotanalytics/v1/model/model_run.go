package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Run struct {

	// 作业运行ID。
	RunId string `json:"run_id"`

	// 作业ID。
	JobId *string `json:"job_id,omitempty"`

	// 作业名称。
	JobName string `json:"job_name"`

	// 作业类型。
	JobType RunJobType `json:"job_type"`

	// 作业开始的时间。时间格式为ISO日期时间格式yyyy-MM-dd'T'HH:mm:ss'Z'
	StartTime string `json:"start_time"`

	// 作业运行时长，单位毫秒。
	Duration int64 `json:"duration"`

	// 此作业的当前状态，包含提交（LAUNCHING）、运行中（RUNNING）、完成（FINISHED）、失败（FAILED）、取消（CANCELLED）。
	Status string `json:"status"`

	// 是否定时作业。
	IsScheduleJob *bool `json:"is_schedule_job,omitempty"`

	// 计算资源名称。
	ComputingResourceName *string `json:"computing_resource_name,omitempty"`

	SqlJob *SqlJobRun `json:"sql_job,omitempty"`
}

func (o Run) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Run struct{}"
	}

	return strings.Join([]string{"Run", string(data)}, " ")
}

type RunJobType struct {
	value string
}

type RunJobTypeEnum struct {
	SQL_JOB RunJobType
}

func GetRunJobTypeEnum() RunJobTypeEnum {
	return RunJobTypeEnum{
		SQL_JOB: RunJobType{
			value: "SqlJob",
		},
	}
}

func (c RunJobType) Value() string {
	return c.value
}

func (c RunJobType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RunJobType) UnmarshalJSON(b []byte) error {
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
