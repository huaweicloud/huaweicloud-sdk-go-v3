package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// NodeExecution 节点执行详情
type NodeExecution struct {

	// 流程节点执行状态
	Status *NodeExecutionStatus `json:"status,omitempty"`

	// 函数执行时的入参
	Input *interface{} `json:"input,omitempty"`

	// 函数执行结果
	Output *interface{} `json:"output,omitempty"`

	// 节点启动时间，UTC毫秒时间戳格式
	BeginTime *int64 `json:"begin_time,omitempty"`

	// 节点结束时间，UTC毫秒时间戳格式
	EndTime *int64 `json:"end_time,omitempty"`

	// 节点错误信息，仅在节点出错时非空
	ErrorMessage *interface{} `json:"error_message,omitempty"`

	// 流程节点请求ID
	RequestId *string `json:"request_id,omitempty"`
}

func (o NodeExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeExecution struct{}"
	}

	return strings.Join([]string{"NodeExecution", string(data)}, " ")
}

type NodeExecutionStatus struct {
	value string
}

type NodeExecutionStatusEnum struct {
	SUCCESS NodeExecutionStatus
	FAIL    NodeExecutionStatus
	RUNNING NodeExecutionStatus
	TIMEOUT NodeExecutionStatus
	CANCEL  NodeExecutionStatus
}

func GetNodeExecutionStatusEnum() NodeExecutionStatusEnum {
	return NodeExecutionStatusEnum{
		SUCCESS: NodeExecutionStatus{
			value: "success",
		},
		FAIL: NodeExecutionStatus{
			value: "fail",
		},
		RUNNING: NodeExecutionStatus{
			value: "running",
		},
		TIMEOUT: NodeExecutionStatus{
			value: "timeout",
		},
		CANCEL: NodeExecutionStatus{
			value: "cancel",
		},
	}
}

func (c NodeExecutionStatus) Value() string {
	return c.value
}

func (c NodeExecutionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeExecutionStatus) UnmarshalJSON(b []byte) error {
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
