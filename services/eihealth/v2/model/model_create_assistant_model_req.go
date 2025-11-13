package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAssistantModelReq **参数解释**： 助手模型创建请求体。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type CreateAssistantModelReq struct {

	// **参数解释**： 服务名称。 **约束限制**： 不涉及 **取值范围**： 支持中英文、数字及 ._-，仅支持中英文、数字开头结尾，长度2-64。 **默认取值**： 不涉及
	ServiceName string `json:"service_name"`

	// **参数解释**： 模型名称。 **约束限制**： 不涉及 **取值范围**： 支持中英文、数字及 ._-，仅支持中英文、数字开头结尾，长度2-64。 **默认取值**： 不涉及
	Name string `json:"name"`

	// **参数解释**： 模型类型。 **约束限制**： 不涉及 **取值范围**： * chat：文本对话模型 * embedding：嵌入模型 **默认取值**： 不涉及
	Type string `json:"type"`

	// **参数解释**： 模型服务API地址。 **约束限制**： 不涉及 **取值范围**： 以http://或https://开头的有效API地址。 **默认取值**： 不涉及
	ModelServiceApi string `json:"model_service_api"`

	// **参数解释**： 模型描述。 **约束限制**： 不涉及 **取值范围**： 字符长度为[0-1000]。 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否支持工具调用。 **约束限制**： 不涉及 **取值范围**： * true：支持工具调用 * false：不支持工具调用 **默认取值**： 不涉及
	FunctionCall bool `json:"function_call"`

	// **参数解释**： 是否支持思维链。 **约束限制**： 不涉及 **取值范围**： * true：支持思维链 * false：不支持思维链 **默认取值**： 不涉及
	ChainOfThought bool `json:"chain_of_thought"`
}

func (o CreateAssistantModelReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAssistantModelReq struct{}"
	}

	return strings.Join([]string{"CreateAssistantModelReq", string(data)}, " ")
}
