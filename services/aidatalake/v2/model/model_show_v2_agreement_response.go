package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowV2AgreementResponse Response Object
type ShowV2AgreementResponse struct {

	// **参数解释**：租户协议列表。
	Agreements *[]TenantV2Agreement `json:"agreements,omitempty"`

	// **参数解释**：租户签署的协议总数。 **取值范围**：0~1000。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowV2AgreementResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowV2AgreementResponse struct{}"
	}

	return strings.Join([]string{"ShowV2AgreementResponse", string(data)}, " ")
}
