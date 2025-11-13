package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllAssistantModelsRequest Request Object
type ListAllAssistantModelsRequest struct {

	// **参数解释**： 是否支持思维链。 **约束限制**： 不涉及 **取值范围**： * true：支持思维链 * false：不支持思维链 **默认取值**： 不涉及
	ChainOfThought *bool `json:"chain_of_thought,omitempty"`

	// **参数解释**： 是否支持工具调用。 **约束限制**： 不涉及 **取值范围**： * true：支持工具调用 * false：不支持工具调用 **默认取值**： 不涉及
	FunctionOfCall *bool `json:"function_of_call,omitempty"`

	// **参数解释**： 模型服务状态。 **约束限制**： 不涉及 **取值范围**： * available：已接入 * unavailable：未接入 **默认取值**： 不涉及
	Status *string `json:"status,omitempty"`

	// **参数解释**： 模型类型。 **约束限制**： 不涉及 **取值范围**： * chat：文本对话模型 * embedding：嵌入式模型 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`
}

func (o ListAllAssistantModelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllAssistantModelsRequest struct{}"
	}

	return strings.Join([]string{"ListAllAssistantModelsRequest", string(data)}, " ")
}
