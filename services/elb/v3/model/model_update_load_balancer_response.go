package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateLoadBalancerResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`
	// 负载均衡器的id（包周期场景返回该字段）

	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
	// 订单号（包周期规格变更场景返回该字段）

	OrderId *string `json:"order_id,omitempty"`
	// 请求ID。  注：自动生成 。

	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerResponse", string(data)}, " ")
}
