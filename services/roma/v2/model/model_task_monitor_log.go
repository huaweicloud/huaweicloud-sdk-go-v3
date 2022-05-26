package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 任务监控日志
type TaskMonitorLog struct {

	// 单次任务执行的跟踪ID
	Id *string `json:"id,omitempty"`

	// 本次执行启动时间，格式timestamp(ms)，使用UTC时区
	StartTime *int64 `json:"start_time,omitempty"`

	// 计划执行时间，格式timestamp(ms)，使用UTC时区
	DispatchTime *int64 `json:"dispatch_time,omitempty"`

	// 写入结束时间，格式timestamp(ms)，使用UTC时区
	EndTime *int64 `json:"end_time,omitempty"`

	// 任务本次执行状态，允许如下值：UNSTARTED-未启动, WAITING-等待调度中, RUNNING-执行中, SUCCESS-执行成功, CANCELLED-任务取消, ERROR-执行异常
	ExecuteStatus *TaskMonitorLogExecuteStatus `json:"execute_status,omitempty"`

	// 标识本次任务执行到哪一个阶段，允许如下值：ADAPTER-任务处于初始化阶段, READER-任务正在执行Reader读操作, WRITER-任务正在执行Writer写操作
	Position *TaskMonitorLogPosition `json:"position,omitempty"`

	// 任务本次执行当前阶段的状态，允许如下值：NORMAL-正在运行, NODE_END-本节点正常结束, RUNTIME_CANCEL-任务被取消, TASK_END-本任务正常结束, RUNTIME_ERR-运行时异常, INTERNAL_ERR-内部程序异常
	PositionStatus *TaskMonitorLogPositionStatus `json:"position_status,omitempty"`

	// 本次任务执行详细状态，使用状态码的形式</br> 状态码划分规则：reader端 100 ~ 499，writer端 500 ~ 899，其他900 ~ </br> 当前状态码如下：</br> 16-被强制取消</br> 99-任务开始</br> 100-Reader 任务开始</br> 101-Reader 任务结束</br> 102-正在读取数据</br> 103-读端数据源端异常</br> 104-读取数据结束</br> 105-读取数据为0</br> 106-读任务强制取消</br> 107-在reader plugin中，任务发生了中断</br> 108-读任务恢复运行</br> 500-Writer 任务开始</br> 501-Writer 任务结束</br> 502-正在数据写入</br> 503-目标端异常</br> 504-数据写入结束</br> 505-写任务强制取消</br> 506-在writer plugin中，任务发生了中断</br> 507-写任务恢复运行</br> 900-接收到调度请求</br> 901-任务运行结束</br> 902-任务已运行结束，正在进行数据完整性校验</br> 903-输出数据完整性校验结果</br> 904-经过数据完整性校验，发现有数据缺失，正在进行数据补偿</br> 905-输出数据补偿结果</br> 906-读取任务正在在排队中（平台资源）</br> 907-读取任务被拒绝执行，因为上一次调度还没有结束</br> 908-写入任务正在在排队中（平台资源）</br> 909-写入任务被拒绝执行，因为上一次调度还没有结束</br> 911-读取任务没有被正常开启，请检查网络是否通畅，参数是否正确</br> 912-写入任务没有被正常开启，请检查网络是否通畅，参数是否正确</br> 913-任务调度请求失败</br> 914-任务被拒绝执行，因为上一次调度还没有结束</br> 915-任务不正常运行</br> 916-任务日志上报异常</br>
	Status *int32 `json:"status,omitempty"`

	// 异常数据条数
	DirtyDataCount *int32 `json:"dirty_data_count,omitempty"`

	// 成功数据条数
	DataCount *int32 `json:"data_count,omitempty"`

	// 成功数据大小，浮点数类型
	DataSize float32 `json:"data_size,omitempty"`

	// 成功数据大小的计量单位
	DataSizeUnit *TaskMonitorLogDataSizeUnit `json:"data_size_unit,omitempty"`

	// 执行时长，单位：ms
	SpendTime *int64 `json:"spend_time,omitempty"`

	// 读取执行时长，单位：ms，只有在定时任务时存在该属性
	ReadSpendTime *int64 `json:"read_spend_time,omitempty"`

	// 写入执行时长，单位：ms
	WriteSpendTime *int64 `json:"write_spend_time,omitempty"`

	// 本次执行结果简要信息
	Remarks *string `json:"remarks,omitempty"`

	// 本次执行详细轨迹信息
	DetailLogs *[]TaskMonitorDetailLog `json:"detail_logs,omitempty"`
}

func (o TaskMonitorLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskMonitorLog struct{}"
	}

	return strings.Join([]string{"TaskMonitorLog", string(data)}, " ")
}

type TaskMonitorLogExecuteStatus struct {
	value string
}

type TaskMonitorLogExecuteStatusEnum struct {
	UNSTARTED TaskMonitorLogExecuteStatus
	WAITING   TaskMonitorLogExecuteStatus
	RUNNING   TaskMonitorLogExecuteStatus
	SUCCESS   TaskMonitorLogExecuteStatus
	CANCELLED TaskMonitorLogExecuteStatus
	ERROR     TaskMonitorLogExecuteStatus
}

func GetTaskMonitorLogExecuteStatusEnum() TaskMonitorLogExecuteStatusEnum {
	return TaskMonitorLogExecuteStatusEnum{
		UNSTARTED: TaskMonitorLogExecuteStatus{
			value: "UNSTARTED",
		},
		WAITING: TaskMonitorLogExecuteStatus{
			value: "WAITING",
		},
		RUNNING: TaskMonitorLogExecuteStatus{
			value: "RUNNING",
		},
		SUCCESS: TaskMonitorLogExecuteStatus{
			value: "SUCCESS",
		},
		CANCELLED: TaskMonitorLogExecuteStatus{
			value: "CANCELLED",
		},
		ERROR: TaskMonitorLogExecuteStatus{
			value: "ERROR",
		},
	}
}

func (c TaskMonitorLogExecuteStatus) Value() string {
	return c.value
}

func (c TaskMonitorLogExecuteStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorLogExecuteStatus) UnmarshalJSON(b []byte) error {
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

type TaskMonitorLogPosition struct {
	value string
}

type TaskMonitorLogPositionEnum struct {
	ADAPTER TaskMonitorLogPosition
	READER  TaskMonitorLogPosition
	WRITER  TaskMonitorLogPosition
}

func GetTaskMonitorLogPositionEnum() TaskMonitorLogPositionEnum {
	return TaskMonitorLogPositionEnum{
		ADAPTER: TaskMonitorLogPosition{
			value: "ADAPTER",
		},
		READER: TaskMonitorLogPosition{
			value: "READER",
		},
		WRITER: TaskMonitorLogPosition{
			value: "WRITER",
		},
	}
}

func (c TaskMonitorLogPosition) Value() string {
	return c.value
}

func (c TaskMonitorLogPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorLogPosition) UnmarshalJSON(b []byte) error {
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

type TaskMonitorLogPositionStatus struct {
	value string
}

type TaskMonitorLogPositionStatusEnum struct {
	NORMAL         TaskMonitorLogPositionStatus
	NODE_END       TaskMonitorLogPositionStatus
	RUNTIME_CANCEL TaskMonitorLogPositionStatus
	TASK_END       TaskMonitorLogPositionStatus
	RUNTIME_ERR    TaskMonitorLogPositionStatus
	INTERNAL_ERR   TaskMonitorLogPositionStatus
}

func GetTaskMonitorLogPositionStatusEnum() TaskMonitorLogPositionStatusEnum {
	return TaskMonitorLogPositionStatusEnum{
		NORMAL: TaskMonitorLogPositionStatus{
			value: "NORMAL",
		},
		NODE_END: TaskMonitorLogPositionStatus{
			value: "NODE_END",
		},
		RUNTIME_CANCEL: TaskMonitorLogPositionStatus{
			value: "RUNTIME_CANCEL",
		},
		TASK_END: TaskMonitorLogPositionStatus{
			value: "TASK_END",
		},
		RUNTIME_ERR: TaskMonitorLogPositionStatus{
			value: "RUNTIME_ERR",
		},
		INTERNAL_ERR: TaskMonitorLogPositionStatus{
			value: "INTERNAL_ERR",
		},
	}
}

func (c TaskMonitorLogPositionStatus) Value() string {
	return c.value
}

func (c TaskMonitorLogPositionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorLogPositionStatus) UnmarshalJSON(b []byte) error {
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

type TaskMonitorLogDataSizeUnit struct {
	value string
}

type TaskMonitorLogDataSizeUnitEnum struct {
	B  TaskMonitorLogDataSizeUnit
	KB TaskMonitorLogDataSizeUnit
	MB TaskMonitorLogDataSizeUnit
	GB TaskMonitorLogDataSizeUnit
	TB TaskMonitorLogDataSizeUnit
}

func GetTaskMonitorLogDataSizeUnitEnum() TaskMonitorLogDataSizeUnitEnum {
	return TaskMonitorLogDataSizeUnitEnum{
		B: TaskMonitorLogDataSizeUnit{
			value: "B",
		},
		KB: TaskMonitorLogDataSizeUnit{
			value: "KB",
		},
		MB: TaskMonitorLogDataSizeUnit{
			value: "MB",
		},
		GB: TaskMonitorLogDataSizeUnit{
			value: "GB",
		},
		TB: TaskMonitorLogDataSizeUnit{
			value: "TB",
		},
	}
}

func (c TaskMonitorLogDataSizeUnit) Value() string {
	return c.value
}

func (c TaskMonitorLogDataSizeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorLogDataSizeUnit) UnmarshalJSON(b []byte) error {
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
