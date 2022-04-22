package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowLoadbalancerTagsRequest struct {

	// 负载均衡器所在的项目id
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancerTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancerTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancerTagsRequest", string(data)}, " ")
}
