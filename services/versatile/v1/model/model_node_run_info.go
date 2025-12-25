package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type NodeRunInfo struct {

	// 工作流id
	AgentId *string `json:"agent_id,omitempty"`

	// 节点id
	NodeId *string `json:"node_id,omitempty"`

	// 父节点id
	ParentNodeId *string `json:"parent_node_id,omitempty"`

	// 工作流节点状态
	NodeStatus *NodeRunInfoNodeStatus `json:"node_status,omitempty"`

	// 父工作流节点id
	ParentWorkflowId *string `json:"parent_workflow_id,omitempty"`

	// 循环次数
	LoopIndex *int32 `json:"loop_index,omitempty"`

	// 上层循环节点id
	LoopNodeId *string `json:"loop_node_id,omitempty"`

	Status *Status `json:"status,omitempty"`

	// 节点名称
	NodeName *string `json:"node_name,omitempty"`

	// 节点类型
	NodeType *string `json:"node_type,omitempty"`

	// 错误信息
	ErrorMessage *string `json:"error_message,omitempty"`

	// 输入参数：debug才设置
	Inputs map[string]interface{} `json:"inputs,omitempty"`

	// 输出参数：debug才设置
	Outputs map[string]interface{} `json:"outputs,omitempty"`

	// 消息节点或提问器节点的输出
	Messages *[]Message `json:"messages,omitempty"`

	// 节点特有的输出，如：大模型原始回复debug才设置
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// 开始时间，13位时间戳 workflow_started,node_started必填
	StartTime *int64 `json:"start_time,omitempty"`

	// 结束时间，13位时间戳 workflow_finished,node_finished必填
	EndTime *int64 `json:"end_time,omitempty"`

	// 真正启动的时间，13位时间戳
	StartupTime *int64 `json:"startup_time,omitempty"`

	// 预热时间，如大模型节点表示首token时间
	PrefillTime *int64 `json:"prefill_time,omitempty"`

	// execution id
	ExecutionId *string `json:"execution_id,omitempty"`

	InnerError *NodeRunInfoInnerError `json:"inner_error,omitempty"`

	// 记忆变量
	Memory map[string]interface{} `json:"memory,omitempty"`
}

func (o NodeRunInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeRunInfo struct{}"
	}

	return strings.Join([]string{"NodeRunInfo", string(data)}, " ")
}

type NodeRunInfoNodeStatus struct {
	value string
}

type NodeRunInfoNodeStatusEnum struct {
	NODE_STARTED  NodeRunInfoNodeStatus
	NODE_WAIT     NodeRunInfoNodeStatus
	NODE_FINISHED NodeRunInfoNodeStatus
}

func GetNodeRunInfoNodeStatusEnum() NodeRunInfoNodeStatusEnum {
	return NodeRunInfoNodeStatusEnum{
		NODE_STARTED: NodeRunInfoNodeStatus{
			value: "node_started",
		},
		NODE_WAIT: NodeRunInfoNodeStatus{
			value: "node_wait",
		},
		NODE_FINISHED: NodeRunInfoNodeStatus{
			value: "node_finished",
		},
	}
}

func (c NodeRunInfoNodeStatus) Value() string {
	return c.value
}

func (c NodeRunInfoNodeStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *NodeRunInfoNodeStatus) UnmarshalJSON(b []byte) error {
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
