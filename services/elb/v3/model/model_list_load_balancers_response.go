package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoadBalancersResponse Response Object
type ListLoadBalancersResponse struct {

	// Loadbalancer的列表。
	Loadbalancers *[]LoadBalancer `json:"loadbalancers,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLoadBalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadBalancersResponse struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersResponse", string(data)}, " ")
}
