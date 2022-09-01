package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLoadBalancersResponse struct {

	// Loadbalancer的列表。
	Loadbalancers *[]LoadBalancer `json:"loadbalancers,omitempty" xml:"loadbalancers"`

	PageInfo *PageInfo `json:"page_info,omitempty" xml:"page_info"`

	// 请求ID。  注：自动生成 。
	RequestId      *string `json:"request_id,omitempty" xml:"request_id"`
	HttpStatusCode int     `json:"-"`
}

func (o ListLoadBalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadBalancersResponse struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersResponse", string(data)}, " ")
}
