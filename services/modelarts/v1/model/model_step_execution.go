package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StepExecution 单节点执行状态。
type StepExecution struct {

	// 节点的名称，在一个DAG中唯一，1到64位只包含中英文，数字，空格，下划线（_）和中划线（-），并且以中英文开头。
	StepName *string `json:"step_name,omitempty"`

	// Execution执行的运行时长。
	Duration *int32 `json:"duration,omitempty"`

	// 节点的类型。
	Type *StepExecutionType `json:"type,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 节点的状态。枚举值如下： - init：初始化 - wait_inputs：等待输入 - pending：等待 - creating：创建中 - created：创建成功 - create_failed：创建失败 - running：运行中 - stopping：停止中 - stopped：停止 - timeout：超时 - completed：完成 - failed：失败 - hold：持有 - skipped：跳过
	Status *StepExecutionStatus `json:"status,omitempty"`

	// 节点的输入项。
	Inputs *[]JobInput `json:"inputs,omitempty"`

	// 节点的输出项。
	Outputs *[]JobOutput `json:"outputs,omitempty"`

	// 节点的UUID，唯一性标识。
	StepUuid *string `json:"step_uuid,omitempty"`

	// 节点的属性。
	Properties map[string]interface{} `json:"properties,omitempty"`

	// 节点发生的事件。
	Events *[]string `json:"events,omitempty"`

	ErrorInfo *WorkflowErrorInfo `json:"error_info,omitempty"`

	Policy *WorkflowStepExecutionPolicy `json:"policy,omitempty"`

	ConditionsExecution *WorkflowConditionExecution `json:"conditions_execution,omitempty"`

	// 节点标题。
	StepTitle *string `json:"step_title,omitempty"`

	Conditions *[]StepCondition `json:"conditions,omitempty"`
}

func (o StepExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepExecution struct{}"
	}

	return strings.Join([]string{"StepExecution", string(data)}, " ")
}

type StepExecutionType struct {
	value string
}

type StepExecutionTypeEnum struct {
	TRANING_JOB StepExecutionType
}

func GetStepExecutionTypeEnum() StepExecutionTypeEnum {
	return StepExecutionTypeEnum{
		TRANING_JOB: StepExecutionType{
			value: "traning_job",
		},
	}
}

func (c StepExecutionType) Value() string {
	return c.value
}

func (c StepExecutionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepExecutionType) UnmarshalJSON(b []byte) error {
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

type StepExecutionStatus struct {
	value string
}

type StepExecutionStatusEnum struct {
	INIT          StepExecutionStatus
	WAIT_INPUTS   StepExecutionStatus
	PENDING       StepExecutionStatus
	CREATING      StepExecutionStatus
	CREATED       StepExecutionStatus
	CREATE_FAILED StepExecutionStatus
	RUNNING       StepExecutionStatus
	STOPPING      StepExecutionStatus
	STOPPED       StepExecutionStatus
	TIMEOUT       StepExecutionStatus
	COMPLETED     StepExecutionStatus
	FAILED        StepExecutionStatus
	HOLD          StepExecutionStatus
	SKIPPED       StepExecutionStatus
}

func GetStepExecutionStatusEnum() StepExecutionStatusEnum {
	return StepExecutionStatusEnum{
		INIT: StepExecutionStatus{
			value: "init",
		},
		WAIT_INPUTS: StepExecutionStatus{
			value: "wait_inputs",
		},
		PENDING: StepExecutionStatus{
			value: "pending",
		},
		CREATING: StepExecutionStatus{
			value: "creating",
		},
		CREATED: StepExecutionStatus{
			value: "created",
		},
		CREATE_FAILED: StepExecutionStatus{
			value: "create_failed",
		},
		RUNNING: StepExecutionStatus{
			value: "running",
		},
		STOPPING: StepExecutionStatus{
			value: "stopping",
		},
		STOPPED: StepExecutionStatus{
			value: "stopped",
		},
		TIMEOUT: StepExecutionStatus{
			value: "timeout",
		},
		COMPLETED: StepExecutionStatus{
			value: "completed",
		},
		FAILED: StepExecutionStatus{
			value: "failed",
		},
		HOLD: StepExecutionStatus{
			value: "hold",
		},
		SKIPPED: StepExecutionStatus{
			value: "skipped",
		},
	}
}

func (c StepExecutionStatus) Value() string {
	return c.value
}

func (c StepExecutionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepExecutionStatus) UnmarshalJSON(b []byte) error {
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
