package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreLoadbalancerRequest Request Object
type RestoreLoadbalancerRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o RestoreLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"RestoreLoadbalancerRequest", string(data)}, " ")
}
