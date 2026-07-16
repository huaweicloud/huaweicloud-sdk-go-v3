package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingQuotaResponse **参数解释**：训练用户配额响应体。
type TrainingQuotaResponse struct {

	// **参数解释**：配额的资源类型，当前支持：job-num，作业的个数配额。 **取值范围**：不涉及。
	Resource *string `json:"resource,omitempty"`

	// **参数解释**：配额个数。 **取值范围**：不涉及。
	Quota *int32 `json:"quota,omitempty"`

	// **参数解释**：已使用的个数。 **取值范围**：不涉及。
	Used *int32 `json:"used,omitempty"`
}

func (o TrainingQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingQuotaResponse struct{}"
	}

	return strings.Join([]string{"TrainingQuotaResponse", string(data)}, " ")
}
