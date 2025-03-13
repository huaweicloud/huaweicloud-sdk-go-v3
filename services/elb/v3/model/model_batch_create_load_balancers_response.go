package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateLoadBalancersResponse Response Object
type BatchCreateLoadBalancersResponse struct {

	// 批创负载均衡器ID（UUID）的列表。
	LoadbalancerIds *[]string `json:"loadbalancer_ids,omitempty"`

	// 批量创建负载均衡器的job ID。
	JobId *string `json:"job_id,omitempty"`

	// 订单号[（只有批量创建包周期实例的场景返回该字段）](tag:hws)  [不支持该字段，请勿使用](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,srg,fcs,fcs_vm,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf,ct)
	OrderId *string `json:"order_id,omitempty"`

	// 请求ID。 注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o BatchCreateLoadBalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadBalancersResponse struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadBalancersResponse", string(data)}, " ")
}
