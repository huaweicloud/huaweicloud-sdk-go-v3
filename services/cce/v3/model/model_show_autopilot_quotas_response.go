package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutopilotQuotasResponse Response Object
type ShowAutopilotQuotasResponse struct {

	// **参数解释：** 资源配额列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Quotas         *[]QuotaResource `json:"quotas,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowAutopilotQuotasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutopilotQuotasResponse struct{}"
	}

	return strings.Join([]string{"ShowAutopilotQuotasResponse", string(data)}, " ")
}
