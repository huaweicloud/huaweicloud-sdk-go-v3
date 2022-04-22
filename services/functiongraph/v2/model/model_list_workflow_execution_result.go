package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 函数流执行概要信息
type ListWorkflowExecutionResult struct {

	// 流程定义ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 唯一标识ID，流程URN
	WorkflowUrn *string `json:"workflow_urn,omitempty"`

	// 流程执行实例ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 流程实例执行状态
	Status *ListWorkflowExecutionResultStatus `json:"status,omitempty"`

	// 流程实例创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	EndTime *string `json:"end_time,omitempty"`

	// 流程实例上次更新时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 流程实例创建者
	CreatedBy *string `json:"created_by,omitempty"`
}

func (o ListWorkflowExecutionResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowExecutionResult struct{}"
	}

	return strings.Join([]string{"ListWorkflowExecutionResult", string(data)}, " ")
}

type ListWorkflowExecutionResultStatus struct {
	value string
}

type ListWorkflowExecutionResultStatusEnum struct {
	SUCCESS ListWorkflowExecutionResultStatus
	FAIL    ListWorkflowExecutionResultStatus
	RUNNING ListWorkflowExecutionResultStatus
	TIMEOUT ListWorkflowExecutionResultStatus
	CANCEL  ListWorkflowExecutionResultStatus
}

func GetListWorkflowExecutionResultStatusEnum() ListWorkflowExecutionResultStatusEnum {
	return ListWorkflowExecutionResultStatusEnum{
		SUCCESS: ListWorkflowExecutionResultStatus{
			value: "success",
		},
		FAIL: ListWorkflowExecutionResultStatus{
			value: "fail",
		},
		RUNNING: ListWorkflowExecutionResultStatus{
			value: "running",
		},
		TIMEOUT: ListWorkflowExecutionResultStatus{
			value: "timeout",
		},
		CANCEL: ListWorkflowExecutionResultStatus{
			value: "cancel",
		},
	}
}

func (c ListWorkflowExecutionResultStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ListWorkflowExecutionResultStatus) UnmarshalJSON(b []byte) error {
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
