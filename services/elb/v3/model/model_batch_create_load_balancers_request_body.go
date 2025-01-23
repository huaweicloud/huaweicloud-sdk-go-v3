package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateLoadBalancersRequestBody This is a auto create Body Object
type BatchCreateLoadBalancersRequestBody struct {
	Loadbalancer *BatchCreateLoadBalancerOption `json:"loadbalancer"`
}

func (o BatchCreateLoadBalancersRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadBalancersRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadBalancersRequestBody", string(data)}, " ")
}
