package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchRemoveAvailableZonesRequest Request Object
type BatchRemoveAvailableZonesRequest struct {

	// 负载均衡器ID。
	LoadbalancerId string `json:"loadbalancer_id"`

	Body *BatchRemoveAvailableZonesRequestBody `json:"body,omitempty"`
}

func (o BatchRemoveAvailableZonesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchRemoveAvailableZonesRequest struct{}"
	}

	return strings.Join([]string{"BatchRemoveAvailableZonesRequest", string(data)}, " ")
}
