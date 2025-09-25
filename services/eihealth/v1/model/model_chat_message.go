package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChatMessage struct {

	// **参数解释**： 对话内容。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Content *string `json:"content,omitempty"`

	// **参数解释**： 深度搜索思维链内容。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ReasoningContent *string `json:"reasoning_content,omitempty"`

	// **参数解释**： 工具调用列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ToolCalls *[]ToolCall `json:"tool_calls,omitempty"`
}

func (o ChatMessage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatMessage struct{}"
	}

	return strings.Join([]string{"ChatMessage", string(data)}, " ")
}
