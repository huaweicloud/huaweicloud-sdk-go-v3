package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadbalancerResponse Response Object
type CreateLoadbalancerResponse struct {
	Loadbalancer *LoadbalancerResp `json:"loadbalancer,omitempty"`

	// 订单号[（包周期场景返回该字段）](tag:hws)
	OrderId *string `json:"order_id,omitempty"`

	// 负载均衡器的ID[（包周期场景返回该字段）](tag:hws)
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLoadbalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerResponse", string(data)}, " ")
}
