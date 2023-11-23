package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddAvailableZonesResponse Response Object
type BatchAddAvailableZonesResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`

	// 请求ID。  注：自动生成 。
	RequestId *string `json:"request_id,omitempty"`

	// 负载均衡器ID，仅在创建包周期场景下返回。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 订单号，仅在创建包周期场景下返回。
	OrderId        *string `json:"order_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchAddAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"BatchAddAvailableZonesResponse", string(data)}, " ")
}
