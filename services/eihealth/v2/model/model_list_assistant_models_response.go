package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAssistantModelsResponse Response Object
type ListAssistantModelsResponse struct {

	// **参数解释**： 供应商模型列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	AssistantModels *[]AssistantModel `json:"assistant_models,omitempty"`

	// **参数解释**： 供应商模型个数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAssistantModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAssistantModelsResponse struct{}"
	}

	return strings.Join([]string{"ListAssistantModelsResponse", string(data)}, " ")
}
