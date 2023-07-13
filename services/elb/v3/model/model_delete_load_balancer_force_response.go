package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerForceResponse Response Object
type DeleteLoadBalancerForceResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerForceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerForceResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerForceResponse", string(data)}, " ")
}
