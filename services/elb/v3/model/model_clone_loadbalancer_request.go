package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CloneLoadbalancerRequest Request Object
type CloneLoadbalancerRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *CloneLoadbalancerRequestBody `json:"body,omitempty"`
}

func (o CloneLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloneLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"CloneLoadbalancerRequest", string(data)}, " ")
}
