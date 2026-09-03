package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ModifyTrainingQuotaRequest **参数解释**：修改训练配额的请求体。
type ModifyTrainingQuotaRequest struct {

	// **参数解释**：用户ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	UserId string `json:"user_id"`

	// **参数解释**：训练作业配额组。
	Quotas []ModifyTrainingQuotaItem `json:"quotas"`
}

func (o ModifyTrainingQuotaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModifyTrainingQuotaRequest struct{}"
	}

	return strings.Join([]string{"ModifyTrainingQuotaRequest", string(data)}, " ")
}
