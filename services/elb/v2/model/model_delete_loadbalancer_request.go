package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteLoadbalancerRequest struct {
	// 负载均衡器id

	LoadbalancerId string `json:"loadbalancer_id"`
	// （不再支持）级联删除负载均衡器

	Cascade *bool `json:"cascade,omitempty"`
}

func (o DeleteLoadbalancerRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLoadbalancerRequest struct{}"
	}

	return strings.Join([]string{"DeleteLoadbalancerRequest", string(data)}, " ")
}
