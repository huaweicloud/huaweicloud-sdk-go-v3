package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AopworkflowInstanceDetail 流程实例详情对象
type AopworkflowInstanceDetail struct {

	// **参数解释**: 流程实例的ID **约束限制**: 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**: 流程实例的名字 **约束限制**: 不涉及
	Name *string `json:"name,omitempty"`

	Workflow *AopworkflowInstanceDetailWorkflow `json:"workflow,omitempty"`

	Dataclass *AopworkflowInstanceDetailDataclass `json:"dataclass,omitempty"`

	Playbook *AopworkflowInstanceDetailPlaybook `json:"playbook,omitempty"`

	// **参数解释**:          触发方式 **取值范围**: - DEBUG   调试触发 - TIMER   定时触发 - EVENT   事件触发 - MANUAL  手动触发
	TriggerType *string `json:"trigger_type,omitempty"`

	// **参数解释**:          流程实例的状态 **取值范围**: - RUNNING   运行中 - FAILED    运行失败 - FINISHED  运行结束 - RETRYING  重试中 - TERMINATING 终止中 - TERMINATED  已终止
	Status *AopworkflowInstanceDetailStatus `json:"status,omitempty"`

	// **参数解释**: 开始时间 **约束限制**: 不涉及
	StartTime *string `json:"start_time,omitempty"`

	// **参数解释**: 结束时间 **约束限制**: 不涉及
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**: 流程实例重试次数 **约束限制**: 不涉及
	RetryCount *int32 `json:"retry_count,omitempty"`

	// **参数解释**: 防线ID **约束限制**: 不涉及
	DefenseId *string `json:"defense_id,omitempty"`

	// **参数解释**: dataobject的ID **约束限制**: 不涉及
	DataobjectId *string `json:"dataobject_id,omitempty"`

	Topology *WorkflowInstanceTopology `json:"topology,omitempty"`
}

func (o AopworkflowInstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AopworkflowInstanceDetail struct{}"
	}

	return strings.Join([]string{"AopworkflowInstanceDetail", string(data)}, " ")
}

type AopworkflowInstanceDetailStatus struct {
	value string
}

type AopworkflowInstanceDetailStatusEnum struct {
	RUNNING     AopworkflowInstanceDetailStatus
	FAILED      AopworkflowInstanceDetailStatus
	FINISHED    AopworkflowInstanceDetailStatus
	RETRYING    AopworkflowInstanceDetailStatus
	TERMINATING AopworkflowInstanceDetailStatus
	TERMINATED  AopworkflowInstanceDetailStatus
}

func GetAopworkflowInstanceDetailStatusEnum() AopworkflowInstanceDetailStatusEnum {
	return AopworkflowInstanceDetailStatusEnum{
		RUNNING: AopworkflowInstanceDetailStatus{
			value: "RUNNING",
		},
		FAILED: AopworkflowInstanceDetailStatus{
			value: "FAILED",
		},
		FINISHED: AopworkflowInstanceDetailStatus{
			value: "FINISHED",
		},
		RETRYING: AopworkflowInstanceDetailStatus{
			value: "RETRYING",
		},
		TERMINATING: AopworkflowInstanceDetailStatus{
			value: "TERMINATING",
		},
		TERMINATED: AopworkflowInstanceDetailStatus{
			value: "TERMINATED",
		},
	}
}

func (c AopworkflowInstanceDetailStatus) Value() string {
	return c.value
}

func (c AopworkflowInstanceDetailStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AopworkflowInstanceDetailStatus) UnmarshalJSON(b []byte) error {
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
