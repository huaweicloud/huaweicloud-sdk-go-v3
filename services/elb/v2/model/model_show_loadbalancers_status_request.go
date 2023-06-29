package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowLoadbalancersStatusRequest Request Object
type ShowLoadbalancersStatusRequest struct {

	// 负载均衡器id
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ShowLoadbalancersStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowLoadbalancersStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowLoadbalancersStatusRequest", string(data)}, " ")
}
