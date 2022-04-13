package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListLoadbalancersResponse struct {
	// 负载均衡器对象列表

	Loadbalancers  *[]LoadbalancerResp `json:"loadbalancers,omitempty"`
	HttpStatusCode int                 `json:"-"`
}

func (o ListLoadbalancersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancersResponse struct{}"
	}

	return strings.Join([]string{"ListLoadbalancersResponse", string(data)}, " ")
}
