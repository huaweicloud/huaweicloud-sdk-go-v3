package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskMonitorDetailLog 任务执行详细跟踪日志
type TaskMonitorDetailLog struct {

	// 任务每次执行步骤产生的唯一ID
	Id *string `json:"id,omitempty"`

	// 当前步骤执行详细状态，使用状态码的形式</br> 状态码划分规则：reader端 100 ~ 499，writer端 500 ~ 899，其他900 ~ </br> 当前状态码如下：</br> 16-被强制取消</br> 99-任务开始</br> 100-Reader 任务开始</br> 101-Reader 任务结束</br> 102-正在读取数据</br> 103-读端数据源端异常</br> 104-读取数据结束</br> 105-读取数据为0</br> 106-读任务强制取消</br> 107-在reader plugin中，任务发生了中断</br> 108-读任务恢复运行</br> 500-Writer 任务开始</br> 501-Writer 任务结束</br> 502-正在数据写入</br> 503-目标端异常</br> 504-数据写入结束</br> 505-写任务强制取消</br> 506-在writer plugin中，任务发生了中断</br> 507-写任务恢复运行</br> 900-接收到调度请求</br> 901-任务运行结束</br> 902-任务已运行结束，正在进行数据完整性校验</br> 903-输出数据完整性校验结果</br> 904-经过数据完整性校验，发现有数据缺失，正在进行数据补偿</br> 905-输出数据补偿结果</br> 906-读取任务正在在排队中（平台资源）</br> 907-读取任务被拒绝执行，因为上一次调度还没有结束</br> 908-写入任务正在在排队中（平台资源）</br> 909-写入任务被拒绝执行，因为上一次调度还没有结束</br> 911-读取任务没有被正常开启，请检查网络是否通畅，参数是否正确</br> 912-写入任务没有被正常开启，请检查网络是否通畅，参数是否正确</br> 913-任务调度请求失败</br> 914-任务被拒绝执行，因为上一次调度还没有结束</br> 915-任务不正常运行</br> 916-任务日志上报异常</br>
	Status *int32 `json:"status,omitempty"`

	// 标识当前步骤属于哪一个阶段，允许如下值：ADAPTER-任务处于初始化阶段, READER-任务正在执行Reader读操作, WRITER-任务正在执行Writer写操作
	Position *TaskMonitorDetailLogPosition `json:"position,omitempty"`

	// 任务当前步骤的状态，允许如下值：NORMAL-正在运行, NODE_END-本节点正常结束, RUNTIME_CANCEL-任务被取消, TASK_END-本任务正常结束, RUNTIME_ERR-运行时异常, INTERNAL_ERR-内部程序异常
	PositionStatus *TaskMonitorDetailLogPositionStatus `json:"position_status,omitempty"`

	// 标识当前步骤属于哪一个FDI插件，如adapter, apireader, rdbwriter等
	Stage *string `json:"stage,omitempty"`

	// 异常数据条数
	DirtyDataCount *int32 `json:"dirty_data_count,omitempty"`

	// 成功数据条数
	DataCount *int32 `json:"data_count,omitempty"`

	// 成功数据大小，浮点数类型
	DataSize float32 `json:"data_size,omitempty"`

	// 成功数据大小的计量单位
	DataSizeUnit *TaskMonitorDetailLogDataSizeUnit `json:"data_size_unit,omitempty"`

	// 执行时长，单位：ms
	SpendTime *int32 `json:"spend_time,omitempty"`

	// 执行详细信息
	Remarks *string `json:"remarks,omitempty"`

	// 本次步骤启动时间，格式timestamp(ms)，使用UTC时区
	StepBeginTime *int64 `json:"step_begin_time,omitempty"`

	// 本次步骤结束时间，格式timestamp(ms)，使用UTC时区
	StepEndTime *int64 `json:"step_end_time,omitempty"`
}

func (o TaskMonitorDetailLog) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskMonitorDetailLog struct{}"
	}

	return strings.Join([]string{"TaskMonitorDetailLog", string(data)}, " ")
}

type TaskMonitorDetailLogPosition struct {
	value string
}

type TaskMonitorDetailLogPositionEnum struct {
	ADAPTER TaskMonitorDetailLogPosition
	READER  TaskMonitorDetailLogPosition
	WRITER  TaskMonitorDetailLogPosition
}

func GetTaskMonitorDetailLogPositionEnum() TaskMonitorDetailLogPositionEnum {
	return TaskMonitorDetailLogPositionEnum{
		ADAPTER: TaskMonitorDetailLogPosition{
			value: "ADAPTER",
		},
		READER: TaskMonitorDetailLogPosition{
			value: "READER",
		},
		WRITER: TaskMonitorDetailLogPosition{
			value: "WRITER",
		},
	}
}

func (c TaskMonitorDetailLogPosition) Value() string {
	return c.value
}

func (c TaskMonitorDetailLogPosition) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorDetailLogPosition) UnmarshalJSON(b []byte) error {
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

type TaskMonitorDetailLogPositionStatus struct {
	value string
}

type TaskMonitorDetailLogPositionStatusEnum struct {
	NORMAL         TaskMonitorDetailLogPositionStatus
	NODE_END       TaskMonitorDetailLogPositionStatus
	RUNTIME_CANCEL TaskMonitorDetailLogPositionStatus
	TASK_END       TaskMonitorDetailLogPositionStatus
	RUNTIME_ERR    TaskMonitorDetailLogPositionStatus
	INTERNAL_ERR   TaskMonitorDetailLogPositionStatus
}

func GetTaskMonitorDetailLogPositionStatusEnum() TaskMonitorDetailLogPositionStatusEnum {
	return TaskMonitorDetailLogPositionStatusEnum{
		NORMAL: TaskMonitorDetailLogPositionStatus{
			value: "NORMAL",
		},
		NODE_END: TaskMonitorDetailLogPositionStatus{
			value: "NODE_END",
		},
		RUNTIME_CANCEL: TaskMonitorDetailLogPositionStatus{
			value: "RUNTIME_CANCEL",
		},
		TASK_END: TaskMonitorDetailLogPositionStatus{
			value: "TASK_END",
		},
		RUNTIME_ERR: TaskMonitorDetailLogPositionStatus{
			value: "RUNTIME_ERR",
		},
		INTERNAL_ERR: TaskMonitorDetailLogPositionStatus{
			value: "INTERNAL_ERR",
		},
	}
}

func (c TaskMonitorDetailLogPositionStatus) Value() string {
	return c.value
}

func (c TaskMonitorDetailLogPositionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorDetailLogPositionStatus) UnmarshalJSON(b []byte) error {
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

type TaskMonitorDetailLogDataSizeUnit struct {
	value string
}

type TaskMonitorDetailLogDataSizeUnitEnum struct {
	B  TaskMonitorDetailLogDataSizeUnit
	KB TaskMonitorDetailLogDataSizeUnit
	MB TaskMonitorDetailLogDataSizeUnit
	GB TaskMonitorDetailLogDataSizeUnit
	TB TaskMonitorDetailLogDataSizeUnit
}

func GetTaskMonitorDetailLogDataSizeUnitEnum() TaskMonitorDetailLogDataSizeUnitEnum {
	return TaskMonitorDetailLogDataSizeUnitEnum{
		B: TaskMonitorDetailLogDataSizeUnit{
			value: "B",
		},
		KB: TaskMonitorDetailLogDataSizeUnit{
			value: "KB",
		},
		MB: TaskMonitorDetailLogDataSizeUnit{
			value: "MB",
		},
		GB: TaskMonitorDetailLogDataSizeUnit{
			value: "GB",
		},
		TB: TaskMonitorDetailLogDataSizeUnit{
			value: "TB",
		},
	}
}

func (c TaskMonitorDetailLogDataSizeUnit) Value() string {
	return c.value
}

func (c TaskMonitorDetailLogDataSizeUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskMonitorDetailLogDataSizeUnit) UnmarshalJSON(b []byte) error {
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
