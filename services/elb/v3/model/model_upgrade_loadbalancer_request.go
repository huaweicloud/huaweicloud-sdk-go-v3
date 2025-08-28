package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeLoadbalancerRequest Request Object
type UpgradeLoadbalancerRequest struct {

	// **参数解释**：要升级的负载均衡器ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *UpgradeV3RequestBody `json:"body,omitempty"`
}

func (o UpgradeLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"UpgradeLoadbalancerRequest", string(data)}, " ")
}
