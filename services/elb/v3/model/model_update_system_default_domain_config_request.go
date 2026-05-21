package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSystemDefaultDomainConfigRequest Request Object
type UpdateSystemDefaultDomainConfigRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpdateSystemDefaultDomainConfigRequestBody `json:"body,omitempty"`
}

func (o UpdateSystemDefaultDomainConfigRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSystemDefaultDomainConfigRequest struct{}"
	}

	return strings.Join([]string{"UpdateSystemDefaultDomainConfigRequest", string(data)}, " ")
}
