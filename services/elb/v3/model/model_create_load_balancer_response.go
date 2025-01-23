package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateLoadBalancerResponse Response Object
type CreateLoadBalancerResponse struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer,omitempty"`

	// 参数解释：负载均衡器的ID。[（包周期场景返回该字段）](tag:hws)  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,srg,fcs,fcs_vm,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,ct)
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`

	// 参数解释：订单号[（包周期场景返回该字段）](tag:hws)  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,srg,fcs,fcs_vm,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,ct)
	OrderId *string `json:"order_id,omitempty"`

	// 参数解释：请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateLoadBalancerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerResponse struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerResponse", string(data)}, " ")
}
