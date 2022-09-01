package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowLoadbalancerRequest struct {

	// 负载均衡器ID
	LoadbalancerId string `json:"loadbalancer_id" xml:"loadbalancer_id"`
}

func (o ShowLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerRequest", string(data)}, " ")
}
