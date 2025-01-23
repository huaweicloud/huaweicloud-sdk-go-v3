package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZonesRequest Request Object
type ListAvailabilityZonesRequest struct {

	// 参数解释：网络公共边界组。
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 参数解释：负载均衡器ID。
	LoadbalancerId *string `json:"loadbalancer_id,omitempty"`
}

func (o ListAvailabilityZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesRequest struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesRequest", string(data)}, " ")
}
