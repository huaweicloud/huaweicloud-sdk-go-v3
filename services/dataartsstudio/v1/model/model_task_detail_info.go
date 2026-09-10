package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskDetailInfo 任务详细信息。
type TaskDetailInfo struct {

	// 集成任务ID。
	TaskId string `json:"task_id"`

	// 任务名称，长度限制0-57个字符。
	TaskName *string `json:"task_name,omitempty"`

	// DLF作业ID。
	MonitorReportId *string `json:"monitor_report_id,omitempty"`

	// 任务类型。 - FLINK：Flink类型 - SPARK：Spark类型 - DRS：DRS类型
	TaskType *TaskDetailInfoTaskType `json:"task_type,omitempty"`

	// 作业运行状态。 - INITIALIZING：初始化中 - SNAPSHOT：全量阶段 - BINLOG：增量阶段
	RunningStatus *TaskDetailInfoRunningStatus `json:"running_status,omitempty"`

	// 计算作业ID，DLI/CCE/DRS执行的作业ID。
	ExternalJobId *string `json:"external_job_id,omitempty"`

	// 源端类型。
	SourceType *string `json:"source_type,omitempty"`

	// 目的端类型。
	TargetType *string `json:"target_type,omitempty"`

	// MRS Flink作业trackingUrl。
	TrackingUrl *string `json:"tracking_url,omitempty"`

	// 任务状态。 - EXCEPTION：异常 - STOPPING：停止中 - SUBMITTING：提交中 - RUNNING：运行中 - STOPPED：已停止 - SUCCESS：成功
	State *TaskDetailInfoState `json:"state,omitempty"`

	// 错误信息。
	ErrorMsg *string `json:"error_msg,omitempty"`

	// 任务创建时间，毫秒时间戳。
	CreateTime *int64 `json:"create_time,omitempty"`

	// 任务更新时间，毫秒时间戳。
	UpdateTime *int64 `json:"update_time,omitempty"`
}

func (o TaskDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskDetailInfo struct{}"
	}

	return strings.Join([]string{"TaskDetailInfo", string(data)}, " ")
}

type TaskDetailInfoTaskType struct {
	value string
}

type TaskDetailInfoTaskTypeEnum struct {
	FLINK TaskDetailInfoTaskType
	SPARK TaskDetailInfoTaskType
	DRS   TaskDetailInfoTaskType
}

func GetTaskDetailInfoTaskTypeEnum() TaskDetailInfoTaskTypeEnum {
	return TaskDetailInfoTaskTypeEnum{
		FLINK: TaskDetailInfoTaskType{
			value: "FLINK",
		},
		SPARK: TaskDetailInfoTaskType{
			value: "SPARK",
		},
		DRS: TaskDetailInfoTaskType{
			value: "DRS",
		},
	}
}

func (c TaskDetailInfoTaskType) Value() string {
	return c.value
}

func (c TaskDetailInfoTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskDetailInfoTaskType) UnmarshalJSON(b []byte) error {
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

type TaskDetailInfoRunningStatus struct {
	value string
}

type TaskDetailInfoRunningStatusEnum struct {
	INITIALIZING TaskDetailInfoRunningStatus
	SNAPSHOT     TaskDetailInfoRunningStatus
	BINLOG       TaskDetailInfoRunningStatus
}

func GetTaskDetailInfoRunningStatusEnum() TaskDetailInfoRunningStatusEnum {
	return TaskDetailInfoRunningStatusEnum{
		INITIALIZING: TaskDetailInfoRunningStatus{
			value: "INITIALIZING",
		},
		SNAPSHOT: TaskDetailInfoRunningStatus{
			value: "SNAPSHOT",
		},
		BINLOG: TaskDetailInfoRunningStatus{
			value: "BINLOG",
		},
	}
}

func (c TaskDetailInfoRunningStatus) Value() string {
	return c.value
}

func (c TaskDetailInfoRunningStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskDetailInfoRunningStatus) UnmarshalJSON(b []byte) error {
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

type TaskDetailInfoState struct {
	value string
}

type TaskDetailInfoStateEnum struct {
	EXCEPTION  TaskDetailInfoState
	STOPPING   TaskDetailInfoState
	SUBMITTING TaskDetailInfoState
	RUNNING    TaskDetailInfoState
	STOPPED    TaskDetailInfoState
	SUCCESS    TaskDetailInfoState
}

func GetTaskDetailInfoStateEnum() TaskDetailInfoStateEnum {
	return TaskDetailInfoStateEnum{
		EXCEPTION: TaskDetailInfoState{
			value: "EXCEPTION",
		},
		STOPPING: TaskDetailInfoState{
			value: "STOPPING",
		},
		SUBMITTING: TaskDetailInfoState{
			value: "SUBMITTING",
		},
		RUNNING: TaskDetailInfoState{
			value: "RUNNING",
		},
		STOPPED: TaskDetailInfoState{
			value: "STOPPED",
		},
		SUCCESS: TaskDetailInfoState{
			value: "SUCCESS",
		},
	}
}

func (c TaskDetailInfoState) Value() string {
	return c.value
}

func (c TaskDetailInfoState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskDetailInfoState) UnmarshalJSON(b []byte) error {
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
