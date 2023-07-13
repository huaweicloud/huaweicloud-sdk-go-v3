package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlowExecutionBrief 函数流执行概要信息
type FlowExecutionBrief struct {

	// 流程定义ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest
	WorkflowUrn *string `json:"workflow_urn,omitempty"`

	// 流程执行实例ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 流程实例执行状态
	Status *FlowExecutionBriefStatus `json:"status,omitempty"`

	// 流程实例创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	EndTime *string `json:"end_time,omitempty"`

	// 流程实例上次更新时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 流程实例创建者
	CreatedBy *string `json:"created_by,omitempty"`
}

func (o FlowExecutionBrief) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlowExecutionBrief struct{}"
	}

	return strings.Join([]string{"FlowExecutionBrief", string(data)}, " ")
}

type FlowExecutionBriefStatus struct {
	value string
}

type FlowExecutionBriefStatusEnum struct {
	SUCCESS FlowExecutionBriefStatus
	FAIL    FlowExecutionBriefStatus
	RUNNING FlowExecutionBriefStatus
	TIMEOUT FlowExecutionBriefStatus
	CANCEL  FlowExecutionBriefStatus
}

func GetFlowExecutionBriefStatusEnum() FlowExecutionBriefStatusEnum {
	return FlowExecutionBriefStatusEnum{
		SUCCESS: FlowExecutionBriefStatus{
			value: "success",
		},
		FAIL: FlowExecutionBriefStatus{
			value: "fail",
		},
		RUNNING: FlowExecutionBriefStatus{
			value: "running",
		},
		TIMEOUT: FlowExecutionBriefStatus{
			value: "timeout",
		},
		CANCEL: FlowExecutionBriefStatus{
			value: "cancel",
		},
	}
}

func (c FlowExecutionBriefStatus) Value() string {
	return c.value
}

func (c FlowExecutionBriefStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlowExecutionBriefStatus) UnmarshalJSON(b []byte) error {
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
