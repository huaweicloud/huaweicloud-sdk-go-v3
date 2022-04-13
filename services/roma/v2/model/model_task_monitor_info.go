package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 任务监控信息
type TaskMonitorInfo struct {
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
	// 任务名称

	TaskName *string `json:"task_name,omitempty"`
	// 任务类型, 只允许两种类型:TIMING-定时任务, REALTIME-实时任务

	TaskType *TaskMonitorInfoTaskType `json:"task_type,omitempty"`
	// 任务状态, 只允许两种类型:0-停止, 1-运行中

	Status *TaskMonitorInfoStatus `json:"status,omitempty"`
	// 任务最近一次执行时间，格式timestamp(ms)，使用UTC时区

	LastExecuteTime *int64 `json:"last_execute_time,omitempty"`
	// 任务是否使用Quartz表达式，只有定时任务才有该属性

	UseQuartzCron *bool `json:"use_quartz_cron,omitempty"`
	// CRON表达式，只有定时任务且use_quartz_cron为true时才有该属性

	Cron *string `json:"cron,omitempty"`
	// 调度周期的单位，如天，小时等，只有定时任务且use_quartz_cron为false时才有该属性，当前仅允许如下类型：MIN-分钟, HOUR-小时, DAY-天, WEEK-周, MON-月

	Period *TaskMonitorInfoPeriod `json:"period,omitempty"`
	// 调度周期，和period字段一起可以确定每隔多长时间调度一次，只有定时任务且use_quartz_cron为false时才有该属性

	DispatchInterval *int32 `json:"dispatch_interval,omitempty"`
	// 标识最近一次任务执行到哪一个阶段，允许如下值：ADAPTER-任务处于初始化阶段, READER-任务正在执行Reader读操作, WRITER-任务正在执行Writer写操作

	Position *TaskMonitorInfoPosition `json:"position,omitempty"`
	// 任务最近一次执行状态，允许如下值：UNSTARTED-未启动, WAITING-等待调度中, RUNNING-执行中, SUCCESS-执行成功, CANCELLED-任务取消, ERROR-执行异常

	ExecuteStatus *TaskMonitorInfoExecuteStatus `json:"execute_status,omitempty"`
	// 任务源端数据源所属应用ID

	SourceAppId *string `json:"source_app_id,omitempty"`
	// 任务源端数据源所属应用名称

	SourceAppName *string `json:"source_app_name,omitempty"`
	// 任务源端数据源所属实例ID

	SourceInstanceId *string `json:"source_instance_id,omitempty"`
	// 任务目标端数据源所属应用ID

	TargetAppId *string `json:"target_app_id,omitempty"`
	// 任务目标端数据源所属应用名称

	TargetAppName *string `json:"target_app_name,omitempty"`
	// 任务目标端数据源所属实例ID

	TargetInstanceId *string `json:"target_instance_id,omitempty"`
	// 任务扩展类型，当前如果是CDC组合任务，该字段为CDC，否则为null

	ExtType *TaskMonitorInfoExtType `json:"ext_type,omitempty"`
	// 任务所属企业项目ID，默认为0

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 任务标签

	TaskTag *string `json:"task_tag,omitempty"`
}

func (o TaskMonitorInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskMonitorInfo struct{}"
	}

	return strings.Join([]string{"TaskMonitorInfo", string(data)}, " ")
}

type TaskMonitorInfoTaskType struct {
	value string
}

type TaskMonitorInfoTaskTypeEnum struct {
	TIMING   TaskMonitorInfoTaskType
	REALTIME TaskMonitorInfoTaskType
}

func GetTaskMonitorInfoTaskTypeEnum() TaskMonitorInfoTaskTypeEnum {
	return TaskMonitorInfoTaskTypeEnum{
		TIMING: TaskMonitorInfoTaskType{
			value: "TIMING",
		},
		REALTIME: TaskMonitorInfoTaskType{
			value: "REALTIME",
		},
	}
}

func (c TaskMonitorInfoTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorInfoTaskType) UnmarshalJSON(b []byte) error {
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

type TaskMonitorInfoStatus struct {
	value int32
}

type TaskMonitorInfoStatusEnum struct {
	E_0 TaskMonitorInfoStatus
	E_1 TaskMonitorInfoStatus
}

func GetTaskMonitorInfoStatusEnum() TaskMonitorInfoStatusEnum {
	return TaskMonitorInfoStatusEnum{
		E_0: TaskMonitorInfoStatus{
			value: 0,
		}, E_1: TaskMonitorInfoStatus{
			value: 1,
		},
	}
}

func (c TaskMonitorInfoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorInfoStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type TaskMonitorInfoPeriod struct {
	value string
}

type TaskMonitorInfoPeriodEnum struct {
	MIN  TaskMonitorInfoPeriod
	HOUR TaskMonitorInfoPeriod
	DAY  TaskMonitorInfoPeriod
	WEEK TaskMonitorInfoPeriod
	MON  TaskMonitorInfoPeriod
}

func GetTaskMonitorInfoPeriodEnum() TaskMonitorInfoPeriodEnum {
	return TaskMonitorInfoPeriodEnum{
		MIN: TaskMonitorInfoPeriod{
			value: "MIN",
		},
		HOUR: TaskMonitorInfoPeriod{
			value: "HOUR",
		},
		DAY: TaskMonitorInfoPeriod{
			value: "DAY",
		},
		WEEK: TaskMonitorInfoPeriod{
			value: "WEEK",
		},
		MON: TaskMonitorInfoPeriod{
			value: "MON",
		},
	}
}

func (c TaskMonitorInfoPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorInfoPeriod) UnmarshalJSON(b []byte) error {
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

type TaskMonitorInfoPosition struct {
	value string
}

type TaskMonitorInfoPositionEnum struct {
	ADAPTER TaskMonitorInfoPosition
	READER  TaskMonitorInfoPosition
	WRITER  TaskMonitorInfoPosition
}

func GetTaskMonitorInfoPositionEnum() TaskMonitorInfoPositionEnum {
	return TaskMonitorInfoPositionEnum{
		ADAPTER: TaskMonitorInfoPosition{
			value: "ADAPTER",
		},
		READER: TaskMonitorInfoPosition{
			value: "READER",
		},
		WRITER: TaskMonitorInfoPosition{
			value: "WRITER",
		},
	}
}

func (c TaskMonitorInfoPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorInfoPosition) UnmarshalJSON(b []byte) error {
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

type TaskMonitorInfoExecuteStatus struct {
	value string
}

type TaskMonitorInfoExecuteStatusEnum struct {
	UNSTARTED TaskMonitorInfoExecuteStatus
	WAITING   TaskMonitorInfoExecuteStatus
	RUNNING   TaskMonitorInfoExecuteStatus
	SUCCESS   TaskMonitorInfoExecuteStatus
	CANCELLED TaskMonitorInfoExecuteStatus
	ERROR     TaskMonitorInfoExecuteStatus
}

func GetTaskMonitorInfoExecuteStatusEnum() TaskMonitorInfoExecuteStatusEnum {
	return TaskMonitorInfoExecuteStatusEnum{
		UNSTARTED: TaskMonitorInfoExecuteStatus{
			value: "UNSTARTED",
		},
		WAITING: TaskMonitorInfoExecuteStatus{
			value: "WAITING",
		},
		RUNNING: TaskMonitorInfoExecuteStatus{
			value: "RUNNING",
		},
		SUCCESS: TaskMonitorInfoExecuteStatus{
			value: "SUCCESS",
		},
		CANCELLED: TaskMonitorInfoExecuteStatus{
			value: "CANCELLED",
		},
		ERROR: TaskMonitorInfoExecuteStatus{
			value: "ERROR",
		},
	}
}

func (c TaskMonitorInfoExecuteStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorInfoExecuteStatus) UnmarshalJSON(b []byte) error {
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

type TaskMonitorInfoExtType struct {
	value string
}

type TaskMonitorInfoExtTypeEnum struct {
	CDC TaskMonitorInfoExtType
}

func GetTaskMonitorInfoExtTypeEnum() TaskMonitorInfoExtTypeEnum {
	return TaskMonitorInfoExtTypeEnum{
		CDC: TaskMonitorInfoExtType{
			value: "CDC",
		},
	}
}

func (c TaskMonitorInfoExtType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorInfoExtType) UnmarshalJSON(b []byte) error {
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
