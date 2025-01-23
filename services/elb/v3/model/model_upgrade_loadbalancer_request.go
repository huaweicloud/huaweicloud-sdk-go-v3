package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpgradeLoadbalancerRequest Request Object
type UpgradeLoadbalancerRequest struct {

	// 负载均衡器ID。
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
