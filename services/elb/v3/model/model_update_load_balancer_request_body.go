package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLoadBalancerRequestBody This is a auto create Body Object
type UpdateLoadBalancerRequestBody struct {
	Loadbalancer *UpdateLoadBalancerOption `json:"loadbalancer"`
}

func (o UpdateLoadBalancerRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerRequestBody", string(data)}, " ")
}
