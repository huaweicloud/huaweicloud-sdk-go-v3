package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerResponse Response Object
type DeleteLoadBalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerResponse", string(data)}, " ")
}
