package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneLoadbalancerRequest Request Object
type CloneLoadbalancerRequest struct {

	// **参数解释**：负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *CloneLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o CloneLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"CloneLoadbalancerRequest", string(data)}, " ")
}
