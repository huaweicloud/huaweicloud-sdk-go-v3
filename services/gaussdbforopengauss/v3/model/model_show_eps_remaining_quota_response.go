package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowEpsRemainingQuotaResponse Response Object
type ShowEpsRemainingQuotaResponse struct {

	// **参数解释**: 剩余企业项目配额组。
	EpsQuotaRemaining *[]EpsRemainingQuotaResult `json:"eps_quota_remaining,omitempty"`

	// **参数解释**: 任务ID。 **取值范围**: 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**: 返回的企业项目个数。 **取值范围**: 不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowEpsRemainingQuotaResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowEpsRemainingQuotaResponse struct{}"
	}

	return strings.Join([]string{"ShowEpsRemainingQuotaResponse", string(data)}, " ")
}
