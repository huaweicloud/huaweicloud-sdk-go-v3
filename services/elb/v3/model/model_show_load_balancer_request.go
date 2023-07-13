package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadBalancerRequest Request Object
type ShowLoadBalancerRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadBalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadBalancerRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadBalancerRequest", string(data)}, " ")
}
