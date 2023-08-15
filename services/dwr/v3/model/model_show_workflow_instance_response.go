package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowWorkflowInstanceResponse Response Object
type ShowWorkflowInstanceResponse struct {

	// 流程定义ID  最小长度：1  最大长度：64
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 函数工作流URN, 格式为： urn:fss:<region_id>:<project_id>:workflow:<package>:<workflow_name>:<version> 注意： package当前只支持default version当前只支持latest
	WorkflowUrn *string `json:"workflow_urn,omitempty"`

	// 流程执行实例ID  最小长度：1  最大长度：64
	ExecutionId *string `json:"execution_id,omitempty"`

	// 流程实例执行状态  最小长度：1  最大长度：32  枚举值：  success  fail  running  timeout  cancel
	Status *ShowWorkflowInstanceResponseStatus `json:"status,omitempty"`

	// 函数执行时需要的Header。
	Headers *interface{} `json:"headers,omitempty"`

	// 函数执行时的入参
	Input *interface{} `json:"input,omitempty"`

	// 函数执行结果
	Output *interface{} `json:"output,omitempty"`

	// 流程实例创建时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间  最小长度：0  最大长度：64
	BeginTime *string `json:"begin_time,omitempty"`

	// 流程实例结束时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间  最小长度：0  最大长度：64
	EndTime *string `json:"end_time,omitempty"`

	// 流程实例上次更新时间，格式：yyyy-MM-ddTHH:mm:ssZ，UTC时间  最小长度：0  最大长度：64
	LastUpdateTime *string `json:"last_update_time,omitempty"`

	// 流程实例创建者  最小长度：1  最大长度：32
	CreatedBy *string `json:"created_by,omitempty"`

	// 节点执行信息
	NodeExecutionDetails *[]NodeExecution `json:"node_execution_details,omitempty"`

	XRequestId *string `json:"x-request-id,omitempty"`

	Connection *string `json:"Connection,omitempty"`

	ContentLength *string `json:"Content-Length,omitempty"`

	Date           *string `json:"Date,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWorkflowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowInstanceResponse", string(data)}, " ")
}

type ShowWorkflowInstanceResponseStatus struct {
	value string
}

type ShowWorkflowInstanceResponseStatusEnum struct {
	SUCCESS ShowWorkflowInstanceResponseStatus
	FAIL    ShowWorkflowInstanceResponseStatus
	RUNNING ShowWorkflowInstanceResponseStatus
	TIMEOUT ShowWorkflowInstanceResponseStatus
	CANCEL  ShowWorkflowInstanceResponseStatus
}

func GetShowWorkflowInstanceResponseStatusEnum() ShowWorkflowInstanceResponseStatusEnum {
	return ShowWorkflowInstanceResponseStatusEnum{
		SUCCESS: ShowWorkflowInstanceResponseStatus{
			value: "success",
		},
		FAIL: ShowWorkflowInstanceResponseStatus{
			value: "fail",
		},
		RUNNING: ShowWorkflowInstanceResponseStatus{
			value: "running",
		},
		TIMEOUT: ShowWorkflowInstanceResponseStatus{
			value: "timeout",
		},
		CANCEL: ShowWorkflowInstanceResponseStatus{
			value: "cancel",
		},
	}
}

func (c ShowWorkflowInstanceResponseStatus) Value() string {
	return c.value
}

func (c ShowWorkflowInstanceResponseStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowWorkflowInstanceResponseStatus) UnmarshalJSON(b []byte) error {
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
