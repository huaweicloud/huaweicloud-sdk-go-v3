package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLoadbalancerRequest Request Object
type DeleteLoadbalancerRequest struct {

	// 负载均衡器id
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o DeleteLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerRequest", string(data)}, " ")
}
