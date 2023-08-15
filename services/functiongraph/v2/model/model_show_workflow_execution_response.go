package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWorkflowExecutionResponse Response Object
type ShowWorkflowExecutionResponse struct {

	// 流程定义ID
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:\\<package\\>:<workflow_name>:\\<version\\> 注意： package当前只支持default version当前只支持latest
	WorkflowUrn *string `json:"workflow_urn,omitempty"`

	// 流程执行实例ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 流程实例执行状态
	Status *ShowWorkflowExecutionResponseStatus `json:"status,omitempty"`

	// 函数执行时需要的Header
	Headers *interface{} `json:"headers,omitempty"`

	// 函数执行时的入参
	Input *interface{} `json:"input,omitempty"`

	// 函数执行结果
	Output *interface{} `json:"output,omitempty"`

	// 流程实例创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	BeginTime *string `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	EndTime *string `json:"end_time,omitempty"`

	// 流程实例上次更新时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 流程实例创建者
	CreatedBy *string `json:"created_by,omitempty"`

	// 节点执行信息
	NodeExecutionDetails *[]NodeExecutionDetail `json:"node_execution_details,omitempty"`
	HttpStatusCode       int                    `json:"-"`
}

func (o ShowWorkflowExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowExecutionResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowExecutionResponse", string(data)}, " ")
}

type ShowWorkflowExecutionResponseStatus struct {
	value string
}

type ShowWorkflowExecutionResponseStatusEnum struct {
	SUCCESS ShowWorkflowExecutionResponseStatus
	FAIL    ShowWorkflowExecutionResponseStatus
	RUNNING ShowWorkflowExecutionResponseStatus
	TIMEOUT ShowWorkflowExecutionResponseStatus
	CANCEL  ShowWorkflowExecutionResponseStatus
}

func GetShowWorkflowExecutionResponseStatusEnum() ShowWorkflowExecutionResponseStatusEnum {
	return ShowWorkflowExecutionResponseStatusEnum{
		SUCCESS: ShowWorkflowExecutionResponseStatus{
			value: "success",
		},
		FAIL: ShowWorkflowExecutionResponseStatus{
			value: "fail",
		},
		RUNNING: ShowWorkflowExecutionResponseStatus{
			value: "running",
		},
		TIMEOUT: ShowWorkflowExecutionResponseStatus{
			value: "timeout",
		},
		CANCEL: ShowWorkflowExecutionResponseStatus{
			value: "cancel",
		},
	}
}

func (c ShowWorkflowExecutionResponseStatus) Value() string {
	return c.value
}

func (c ShowWorkflowExecutionResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWorkflowExecutionResponseStatus) UnmarshalJSON(b []byte) error {
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
