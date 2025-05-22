package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteRecycleLoadBalancerResponse Response Object
type DeleteRecycleLoadBalancerResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRecycleLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteRecycleLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecycleLoadBalancerResponse", string(data)}, " ")
}
