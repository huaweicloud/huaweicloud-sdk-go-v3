package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateLoadBalancerResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty" xml:"loadbalancer"`

	// 负载均衡器的id（包周期场景返回该字段）
	LoadbalancerId *string `json:"loadbalancer_id,omitempty" xml:"loadbalancer_id"`

	// 订单号（包周期场景返回该字段）
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerResponse", string(data)}, " ")
}
