package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunWorkflowRequest Request Object
type RunWorkflowRequest struct {

	// **参数解释**： 工作流ID  **取值范围**： 由英文，数字，“-”，“_”组成，不超过64位字符。
	WorkflowId string `json:"workflow_id"`

	// **参数解释**： 会话ID，历史对话记忆功能基于同一个会话  **取值范围**： 由英文，数字，“-”，“_”组成，不超过64位字符。
	ConversationId string `json:"conversation_id"`

	// **参数解释**： 空间ID，当前资源所属工作空间  **取值范围**： 由英文，数字，“-”，“_”组成，不超过64位字符。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**： 运行的智能体/工作流的发布版本号，默认运行最新发布版本  **取值范围**： 不涉及
	Version *string `json:"version,omitempty"`

	// **参数解释**： 是否流式响应  **取值范围**： true,false（不传默认 true，支持流式）
	Stream *bool `json:"stream,omitempty"`

	Body *WorkflowRunReq `json:"body,omitempty"`
}

func (o RunWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunWorkflowRequest struct{}"
	}

	return strings.Join([]string{"RunWorkflowRequest", string(data)}, " ")
}
