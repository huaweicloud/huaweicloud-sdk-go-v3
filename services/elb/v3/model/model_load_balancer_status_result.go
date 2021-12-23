package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type LoadBalancerStatusResult struct {
	Loadbalancer *LoadBalancerStatus `json:"loadbalancer"`
}

func (o LoadBalancerStatusResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancerStatusResult struct{}"
	}

	return strings.Join([]string{"LoadBalancerStatusResult", string(data)}, " ")
}
