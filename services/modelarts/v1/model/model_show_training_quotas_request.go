package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowTrainingQuotasRequest Request Object
type ShowTrainingQuotasRequest struct {

	// **参数解释**：用户ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：配额的资源类型。 **约束限制**：当前支持的传参：job-num（作业个数的配额）。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Resource *string `json:"resource,omitempty"`
}

func (o ShowTrainingQuotasRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingQuotasRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingQuotasRequest", string(data)}, " ")
}
