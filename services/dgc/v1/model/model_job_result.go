package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobResult struct {

	// 作业名称
	Name string `json:"name"`

	// 作业类型
	JobType string `json:"jobType"`

	// 作业责任人，长度不能超过128个字符
	Owner *string `json:"owner,omitempty"`

	// 作业优先级，取值范围[0, 2]，默认值是0。0代表高优先级，1代表中优先级，2代表低优先级
	Priority *int32 `json:"priority,omitempty"`

	Status JobResultStatus `json:"status"`

	// 作业的创建者
	CreateUser string `json:"createUser"`

	// 作业的创建时间
	CreateTime int64 `json:"createTime"`

	// 作业的启动时间
	StartTime *int64 `json:"startTime,omitempty"`

	// 作业配置的结束时间
	EndTime *int64 `json:"endTime,omitempty"`

	// 作业最近一次运行实例状态，当jobType为BATCH时才有本字段
	LastInstanceStatus *string `json:"lastInstanceStatus,omitempty"`

	// 作业最近一次运行实例运行结束时间，当jobType为BATCH时才有本字段
	LastInstanceEndTime *int64 `json:"lastInstanceEndTime,omitempty"`

	// 作业最后一次更新时间
	LastUpdateTime *int64 `json:"lastUpdateTime,omitempty"`

	// 作业最后一次更新用户
	LastUpdateUser *string `json:"lastUpdateUser,omitempty"`

	// 作业的路径
	Path *string `json:"path,omitempty"`

	// 作业是否为单任务作业
	SingleNodeJobFlag *bool `json:"singleNodeJobFlag,omitempty"`

	// flink作业信息
	FlinkJobInfo *string `json:"flinkJobInfo,omitempty"`

	// 作业监控告警信息
	Alarms *[]JobAlarm `json:"alarms,omitempty"`
}

func (o JobResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobResult struct{}"
	}

	return strings.Join([]string{"JobResult", string(data)}, " ")
}

type JobResultStatus struct {
	value string
}

type JobResultStatusEnum struct {
	STARTING   JobResultStatus
	NORMAL     JobResultStatus
	EXCEPTION  JobResultStatus
	STOPPING   JobResultStatus
	STOPPED    JobResultStatus
	SCHEDULING JobResultStatus
	PAUSED     JobResultStatus
}

func GetJobResultStatusEnum() JobResultStatusEnum {
	return JobResultStatusEnum{
		STARTING: JobResultStatus{
			value: "STARTING",
		},
		NORMAL: JobResultStatus{
			value: "NORMAL",
		},
		EXCEPTION: JobResultStatus{
			value: "EXCEPTION",
		},
		STOPPING: JobResultStatus{
			value: "STOPPING",
		},
		STOPPED: JobResultStatus{
			value: "STOPPED",
		},
		SCHEDULING: JobResultStatus{
			value: "SCHEDULING",
		},
		PAUSED: JobResultStatus{
			value: "PAUSED",
		},
	}
}

func (c JobResultStatus) Value() string {
	return c.value
}

func (c JobResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobResultStatus) UnmarshalJSON(b []byte) error {
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
