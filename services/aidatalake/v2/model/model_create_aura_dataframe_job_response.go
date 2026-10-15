package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAuraDataframeJobResponse Response Object
type CreateAuraDataframeJobResponse struct {

	// **参数解释**：作业ID。  **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**：作业状态。  **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及。
	Status *string `json:"status,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateAuraDataframeJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAuraDataframeJobResponse struct{}"
	}

	return strings.Join([]string{"CreateAuraDataframeJobResponse", string(data)}, " ")
}
