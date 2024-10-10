package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAddAvailableZonesRequest Request Object
type BatchAddAvailableZonesRequest struct {

	// 参数解释：负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchAddAvailableZonesRequestBody `json:"body,omitempty"`
}

func (o BatchAddAvailableZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAddAvailableZonesRequest struct{}"
	}

	return strings.Join([]string{"BatchAddAvailableZonesRequest", string(data)}, " ")
}
