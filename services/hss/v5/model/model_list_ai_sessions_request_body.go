package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListAiSessionsRequestBody struct {

	// **参数解释**： 会话ID **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	SessionId *string `json:"session_id,omitempty"`

	// **参数解释**： 智能体名字 **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	AgentName *string `json:"agent_name,omitempty"`

	// **参数解释**： 智能体ID **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 智能体运行时名字 **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	RuntimeName *string `json:"runtime_name,omitempty"`

	// **参数解释**： 智能体运行时ID **约束限制**： 不涉及 **取值范围**： 字符长度为1-128位 **默认取值**： 不涉及
	RuntimeId *string `json:"runtime_id,omitempty"`

	// **参数解释**： SDK检测状态 **约束限制**： 不涉及 **取值范围**： - open：开启 - close：关闭  **默认取值**： 不涉及
	SdkStatus *string `json:"sdk_status,omitempty"`

	// **参数解释**： 智能体的来源 **约束限制**： 不涉及 **取值范围**： - AgentArts：AgentArts。 - OfficeClaw：OfficeClaw。  **默认取值**： 不涉及
	Source *string `json:"source,omitempty"`
}

func (o ListAiSessionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiSessionsRequestBody struct{}"
	}

	return strings.Join([]string{"ListAiSessionsRequestBody", string(data)}, " ")
}
