package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAuraDataframeJobResponse Response Object
type ListAuraDataframeJobResponse struct {

	// **参数解释**：作业响应列表。 **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及。
	JobList *[]ShowDataframeJobResp `json:"job_list,omitempty"`

	// **参数解释**：作业总数。 **约束限制**：不涉及。 **取值范围**：不涉及 **默认取值**：不涉及。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAuraDataframeJobResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAuraDataframeJobResponse struct{}"
	}

	return strings.Join([]string{"ListAuraDataframeJobResponse", string(data)}, " ")
}
