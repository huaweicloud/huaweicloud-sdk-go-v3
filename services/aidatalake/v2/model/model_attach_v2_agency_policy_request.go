package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachV2AgencyPolicyRequest Request Object
type AttachV2AgencyPolicyRequest struct {

	// **参数解释**：委托名称。 **约束限制**：固定值ai_datalake_admin_trust。 **取值范围**：不涉及。 **默认取值**：不涉及。
	AgencyName string `json:"agency_name"`

	Body *AttachV2AgencyPolicyRequestBody `json:"body,omitempty"`
}

func (o AttachV2AgencyPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachV2AgencyPolicyRequest struct{}"
	}

	return strings.Join([]string{"AttachV2AgencyPolicyRequest", string(data)}, " ")
}
