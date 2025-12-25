package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Message struct {

	// 角色
	Role *string `json:"role,omitempty"`

	// 内容
	Content *string `json:"content,omitempty"`

	// 时间戳
	CreateTime *int64 `json:"create_time,omitempty"`

	// 工具名称
	Name *interface{} `json:"name,omitempty"`

	// 工具调用信息
	FunctionCall *interface{} `json:"function_call,omitempty"`

	// 工具调用信息
	ToolCalls *interface{} `json:"tool_calls,omitempty"`

	// 工具调用ID信息
	ToolCallId *interface{} `json:"tool_call_id,omitempty"`

	// 是否开启会话历史
	EnableHistory *bool `json:"enable_history,omitempty"`

	Intent *[]string `json:"intent,omitempty"`

	// 对话ID
	ExecutionId *string `json:"execution_id,omitempty"`

	// 节点ID
	NodeId *string `json:"node_id,omitempty"`

	// agent id
	AgentId *string `json:"agent_id,omitempty"`

	// 评分，-1（点赞）、1（点踩）
	Rating *int32 `json:"rating,omitempty"`

	// 多模态文件列表
	Files *[]interface{} `json:"files,omitempty"`

	Reason *FeedbackReason `json:"reason,omitempty"`
}

func (o Message) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Message struct{}"
	}

	return strings.Join([]string{"Message", string(data)}, " ")
}
