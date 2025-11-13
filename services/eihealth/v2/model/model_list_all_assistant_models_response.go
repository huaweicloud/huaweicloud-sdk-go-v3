package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllAssistantModelsResponse Response Object
type ListAllAssistantModelsResponse struct {

	// **参数解释**： 供应商列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Vendors *[]AssistantModelVendorInfo `json:"vendors,omitempty"`

	// **参数解释**： 供应商个数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAllAssistantModelsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllAssistantModelsResponse struct{}"
	}

	return strings.Join([]string{"ListAllAssistantModelsResponse", string(data)}, " ")
}
