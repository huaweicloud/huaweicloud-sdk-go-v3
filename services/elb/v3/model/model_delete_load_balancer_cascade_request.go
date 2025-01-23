package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerCascadeRequest Request Object
type DeleteLoadBalancerCascadeRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *DeleteLoadBalancerCascadeRequestBody `json:"body,omitempty"`
}

func (o DeleteLoadBalancerCascadeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerCascadeRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerCascadeRequest", string(data)}, " ")
}
