package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLoadbalancerFeatureRequest Request Object
type ListLoadbalancerFeatureRequest struct {

	// ELB实例ID。
	LoadbalancerId string `json:"loadbalancer_id"`
}

func (o ListLoadbalancerFeatureRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadbalancerFeatureRequest struct{}"
	}

	return strings.Join([]string{"ListLoadbalancerFeatureRequest", string(data)}, " ")
}
