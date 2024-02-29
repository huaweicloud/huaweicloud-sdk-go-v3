package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// JobInstance 作业实例信息
type JobInstance struct {

	// 作业名称。如果要查询指定批处理作业的实例列表，jobName就是批处理作业名。如果要查询实时作业下某个节点关联的子作业，jobName格式为[实时作业名称]_[节点名称]。
	JobName *string `json:"job_name,omitempty"`

	// 实例运行状态： - waiting：等待运行 - running：运行中 - success：运行成功 - fail： 运行失败 - running-exception：运行异常 - pause： 暂停 - manual-stop：取消 - skip-by-depend：跳过 - freeze：冻结
	Status *JobInstanceStatus `json:"status,omitempty"`

	// 作业实例运行时日志记录的实例名称, 非作业定义的名称
	JobInstanceName *string `json:"job_instance_name,omitempty"`

	// 作业实例计划执行时间
	PlanTime *int64 `json:"plan_time,omitempty"`

	// 作业实例实际执行开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 作业实例实际执行结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 执行耗时，单位：毫秒
	ExecuteTime *int64 `json:"execute_time,omitempty"`

	// 作业实例id
	InstanceId *int64 `json:"instance_id,omitempty"`

	// 作业实例类型
	InstanceType *int32 `json:"instance_type,omitempty"`

	// 作业提交运行时间
	SubmitTime *int64 `json:"submit_time,omitempty"`

	// 作业id
	JobId *int64 `json:"job_id,omitempty"`

	// 作业实例版本
	Version *int32 `json:"version,omitempty"`

	// 作业实例状态筛选为强制成功
	ForceSuccess *bool `json:"force_success,omitempty"`

	// 作业实例状态筛选为忽略失败
	IgnoreSuccess *bool `json:"ignore_success,omitempty"`
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
	SKIP_BY_DEPEND    JobInstanceStatus
	FREEZE            JobInstanceStatus
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
		SKIP_BY_DEPEND: JobInstanceStatus{
			value: "skip-by-depend",
		},
		FREEZE: JobInstanceStatus{
			value: "freeze",
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
