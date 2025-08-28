package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleBinLoadBalancersResponse Response Object
type ListRecycleBinLoadBalancersResponse struct {

	// **参数解释**：回收站中的弹性负载均衡器实例列表。  **取值范围**：不涉及
	Loadbalancers *[]RecycleLoadBalancer `json:"loadbalancers,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// **参数解释**：请求ID。  **取值范围**：由数字、小写字母和中划线（-）组成的字符串，自动生成。
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRecycleBinLoadBalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleBinLoadBalancersResponse struct{}"
	}

	return strings.Join([]string{"ListRecycleBinLoadBalancersResponse", string(data)}, " ")
}
