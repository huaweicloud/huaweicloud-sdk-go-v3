package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StepExecutionResp 单节点执行状态。
type StepExecutionResp struct {

	// **参数解释**：节点的名称，在一个DAG中唯一。 **取值范围**：不涉及。
	StepName *string `json:"step_name,omitempty"`

	// **参数解释**：执行记录的名称。 **取值范围**：不涉及。
	ExecutionName *string `json:"execution_name,omitempty"`

	// **参数解释**：执行记录与节点的组合名称。 **取值范围**：不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：唯一标识uuid。创建节点执行时，后台自动生成。 **取值范围**：不涉及。
	Uuid *string `json:"uuid,omitempty"`

	// **参数解释**：执行记录的UUID。 **取值范围**：不涉及。
	ExecutionUuid *string `json:"execution_uuid,omitempty"`

	// **参数解释**：Execution执行的创建时间。 **取值范围**：不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：Execution执行的更新时间。 **取值范围**：不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**：Execution执行的运行时长。 **取值范围**：不涉及。
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**：节点的类型。 **取值范围**：枚举值如下: - job：训练 - labeling：标注 - release_dataset：数据集发布 - model：模型发布 - service：服务部署 - mrs_job：MRS作业 - dataset_import：数据集导入 - create_dataset：创建数据集
	Type *StepExecutionRespType `json:"type,omitempty"`

	// **参数解释**：实例ID。 **取值范围**：不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：节点的状态。 **取值范围**：枚举值如下： - init：初始化 - wait_inputs：等待输入 - pending：等待 - creating：创建中 - created：创建成功 - create_failed：创建失败 - running：运行中 - stopping：停止中 - stopped：停止 - timeout：超时 - completed：完成 - failed：失败 - hold：暂停 - skipped：跳过
	Status *StepExecutionRespStatus `json:"status,omitempty"`

	// **参数解释**：节点的输入项。
	Inputs *[]JobInputResp `json:"inputs,omitempty"`

	// **参数解释**：节点的输出项。
	Outputs *[]JobOutputResp `json:"outputs,omitempty"`

	// **参数解释**：节点的UUID，唯一性标识。 **取值范围**：不涉及。
	StepUuid *string `json:"step_uuid,omitempty"`

	// **参数解释**：节点的属性。
	Properties map[string]string `json:"properties,omitempty"`

	// **参数解释**：节点发生的事件。
	Events *[]string `json:"events,omitempty"`

	ErrorInfo *WorkflowErrorInfoResp `json:"error_info,omitempty"`

	Policy *WorkflowStepExecutionPolicyResp `json:"policy,omitempty"`

	ConditionsExecution *WorkflowConditionExecutionResp `json:"conditions_execution,omitempty"`

	// **参数解释**：节点标题。 **取值范围**：不涉及。
	StepTitle *string `json:"step_title,omitempty"`

	// **参数解释**：条件节点执行条件。
	Conditions *[]StepConditionResp `json:"conditions,omitempty"`
}

func (o StepExecutionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StepExecutionResp struct{}"
	}

	return strings.Join([]string{"StepExecutionResp", string(data)}, " ")
}

type StepExecutionRespType struct {
	value string
}

type StepExecutionRespTypeEnum struct {
	TRANING_JOB StepExecutionRespType
}

func GetStepExecutionRespTypeEnum() StepExecutionRespTypeEnum {
	return StepExecutionRespTypeEnum{
		TRANING_JOB: StepExecutionRespType{
			value: "traning_job",
		},
	}
}

func (c StepExecutionRespType) Value() string {
	return c.value
}

func (c StepExecutionRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepExecutionRespType) UnmarshalJSON(b []byte) error {
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

type StepExecutionRespStatus struct {
	value string
}

type StepExecutionRespStatusEnum struct {
	INIT          StepExecutionRespStatus
	WAIT_INPUTS   StepExecutionRespStatus
	PENDING       StepExecutionRespStatus
	CREATING      StepExecutionRespStatus
	CREATED       StepExecutionRespStatus
	CREATE_FAILED StepExecutionRespStatus
	RUNNING       StepExecutionRespStatus
	STOPPING      StepExecutionRespStatus
	STOPPED       StepExecutionRespStatus
	TIMEOUT       StepExecutionRespStatus
	COMPLETED     StepExecutionRespStatus
	FAILED        StepExecutionRespStatus
	HOLD          StepExecutionRespStatus
	SKIPPED       StepExecutionRespStatus
}

func GetStepExecutionRespStatusEnum() StepExecutionRespStatusEnum {
	return StepExecutionRespStatusEnum{
		INIT: StepExecutionRespStatus{
			value: "init",
		},
		WAIT_INPUTS: StepExecutionRespStatus{
			value: "wait_inputs",
		},
		PENDING: StepExecutionRespStatus{
			value: "pending",
		},
		CREATING: StepExecutionRespStatus{
			value: "creating",
		},
		CREATED: StepExecutionRespStatus{
			value: "created",
		},
		CREATE_FAILED: StepExecutionRespStatus{
			value: "create_failed",
		},
		RUNNING: StepExecutionRespStatus{
			value: "running",
		},
		STOPPING: StepExecutionRespStatus{
			value: "stopping",
		},
		STOPPED: StepExecutionRespStatus{
			value: "stopped",
		},
		TIMEOUT: StepExecutionRespStatus{
			value: "timeout",
		},
		COMPLETED: StepExecutionRespStatus{
			value: "completed",
		},
		FAILED: StepExecutionRespStatus{
			value: "failed",
		},
		HOLD: StepExecutionRespStatus{
			value: "hold",
		},
		SKIPPED: StepExecutionRespStatus{
			value: "skipped",
		},
	}
}

func (c StepExecutionRespStatus) Value() string {
	return c.value
}

func (c StepExecutionRespStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StepExecutionRespStatus) UnmarshalJSON(b []byte) error {
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
