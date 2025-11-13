package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssistantModel 供应商模型详情。
type AssistantModel struct {

	// **参数解释**： 模型供应商ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 服务名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ServiceName *string `json:"service_name,omitempty"`

	// **参数解释**： 模型名称。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 模型类型。 **约束限制**： 不涉及 **取值范围**： * CHAT：文本对话模型 * EMBEDDING：嵌入模型 **默认取值**： 不涉及
	Type *string `json:"type,omitempty"`

	// **参数解释**： 模型服务API地址。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ModelServiceApi *string `json:"model_service_api,omitempty"`

	// **参数解释**： 是否支持工具调用。 **约束限制**： 不涉及 **取值范围**： * true：支持工具调用 * false：不支持工具调用 **默认取值**： 不涉及
	FunctionCall *bool `json:"function_call,omitempty"`

	// **参数解释**： 是否支持思维链。 **约束限制**： 不涉及 **取值范围**： * true：支持思维链 * false：不支持思维链 **默认取值**： 不涉及
	ChainOfThought *bool `json:"chain_of_thought,omitempty"`

	// **参数解释**： 模型创建人。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 模型创建人ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreatorId *string `json:"creator_id,omitempty"`

	// **参数解释**： 模型创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 模型修改时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 模型描述。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Description *string `json:"description,omitempty"`
}

func (o AssistantModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssistantModel struct{}"
	}

	return strings.Join([]string{"AssistantModel", string(data)}, " ")
}
