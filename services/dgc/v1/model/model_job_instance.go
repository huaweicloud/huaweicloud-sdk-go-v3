package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type JobInstance struct {

	// 作业名称。如果要查询指定批处理作业的实例列表，jobName就是批处理作业名。如果要查询实时作业下某个节点关联的子作业，jobName格式为[实时作业名称]_[节点名称]。
	JobName *string `json:"jobName,omitempty"`

	// 实例运行状态： - waiting：等待运行 - running：运行中 - success：运行成功 - fail： 运行失败 - running-exception：运行异常 - pause： 暂停 - manual-stop：取消
	Status *JobInstanceStatus `json:"status,omitempty"`

	// 作业实例计划执行时间
	PlanTime *int64 `json:"planTime,omitempty"`

	// 作业实例实际执行开始时间
	StartTime *int64 `json:"startTime,omitempty"`

	// 作业实例实际执行结束时间
	EndTime *int64 `json:"endTime,omitempty"`

	// 执行耗时，单位：毫秒
	ExecuteTime *int64 `json:"executeTime,omitempty"`

	// 作业提交运行时间
	SubmitTime *int64 `json:"submitTime,omitempty"`

	// 作业实例ID
	InstanceId *int64 `json:"instanceId,omitempty"`

	// 作业ID
	JobId *int64 `json:"jobId,omitempty"`

	// 作业实例运行时日志记录的实例名称, 非作业定义的名称
	JobInstanceName *string `json:"jobInstanceName,omitempty"`

	// 作业实例类型
	InstanceType *int32 `json:"instanceType,omitempty"`

	// 作业版本号
	Version *int32 `json:"version,omitempty"`

	// 作业成功状态，是否忽略成功
	IgnoreSuccess *bool `json:"ignoreSuccess,omitempty"`

	// 作业成功状态，是否强制成功
	ForceSuccess *bool `json:"forceSuccess,omitempty"`
}

func (o JobInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobInstance struct{}"
	}

	return strings.Join([]string{"JobInstance", string(data)}, " ")
}

type JobInstanceStatus struct {
	value string
}

type JobInstanceStatusEnum struct {
	WAITING           JobInstanceStatus
	RUNNING           JobInstanceStatus
	SUCCESS           JobInstanceStatus
	FAIL              JobInstanceStatus
	RUNNING_EXCEPTION JobInstanceStatus
	PAUSE             JobInstanceStatus
	MANUAL_STOP       JobInstanceStatus
}

func GetJobInstanceStatusEnum() JobInstanceStatusEnum {
	return JobInstanceStatusEnum{
		WAITING: JobInstanceStatus{
			value: "waiting",
		},
		RUNNING: JobInstanceStatus{
			value: "running",
		},
		SUCCESS: JobInstanceStatus{
			value: "success",
		},
		FAIL: JobInstanceStatus{
			value: "fail",
		},
		RUNNING_EXCEPTION: JobInstanceStatus{
			value: "running-exception",
		},
		PAUSE: JobInstanceStatus{
			value: "pause",
		},
		MANUAL_STOP: JobInstanceStatus{
			value: "manual-stop",
		},
	}
}

func (c JobInstanceStatus) Value() string {
	return c.value
}

func (c JobInstanceStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *JobInstanceStatus) UnmarshalJSON(b []byte) error {
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
