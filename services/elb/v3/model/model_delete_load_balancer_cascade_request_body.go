package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerCascadeRequestBody This is a auto create Body Object
type DeleteLoadBalancerCascadeRequestBody struct {
	Loadbalancer *DeleteLoadBalancerCascadeOption `json:"loadbalancer"`
}

func (o DeleteLoadBalancerCascadeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerCascadeRequestBody struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerCascadeRequestBody", string(data)}, " ")
}
