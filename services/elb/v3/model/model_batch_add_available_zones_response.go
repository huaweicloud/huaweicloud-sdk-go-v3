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

	// 负载均衡器ID[（包周期场景返回该字段）](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,ctc,cmcc,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 订单号[（包周期场景返回该字段）](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,ctc,cmcc,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
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
