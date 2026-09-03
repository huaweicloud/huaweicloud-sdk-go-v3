package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyTrainingQuotasResponse Response Object
type ModifyTrainingQuotasResponse struct {

	// **参数解释**：用户ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**：训练作业配额组。
	Quotas         *[]ModifyTrainingQuotaItem `json:"quotas,omitempty"`
	HttpStatusCode int                        `json:"-"`
}

func (o ModifyTrainingQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTrainingQuotasResponse struct{}"
	}

	return strings.Join([]string{"ModifyTrainingQuotasResponse", string(data)}, " ")
}
