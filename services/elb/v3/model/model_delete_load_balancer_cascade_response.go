package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadBalancerCascadeResponse Response Object
type DeleteLoadBalancerCascadeResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteLoadBalancerCascadeResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadBalancerCascadeResponse struct{}"
	}

	return strings.Join([]string{"DeleteLoadBalancerCascadeResponse", string(data)}, " ")
}
