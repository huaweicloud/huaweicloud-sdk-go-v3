package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowChatResponse Response Object
type ShowChatResponse struct {

	// **参数解释**： 对话ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 对话名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Title *string `json:"title,omitempty"`

	// **参数解释**： 对话别名。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Alias *string `json:"alias,omitempty"`

	// **参数解释**： 对话创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 对话更新时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 绑定的知识库ID列表。 **约束限制**： 绑定的知识库数量范围为[0-5]。 **取值范围**： 不涉及 **默认取值**： 不涉及
	RepoIds *[]string `json:"repo_ids,omitempty"`

	AgentType *AgentType `json:"agent_type,omitempty"`

	AgentRole *AgentRole `json:"agent_role,omitempty"`

	// **参数解释**： 对话问答列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ChatMessages *[]ChatMessageRsp `json:"chat_messages,omitempty"`

	XChatRouteId   *string `json:"X-Chat-Route-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowChatResponse struct{}"
	}

	return strings.Join([]string{"ShowChatResponse", string(data)}, " ")
}
