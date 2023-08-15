package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZonesResponse Response Object
type ListAvailabilityZonesResponse struct {

	// 云桌面支持的可用分区列表。
	AvailabilityZones *[]AvailabilityZoneInfo `json:"availability_zones,omitempty"`

	// 云桌面支持的可用分区列表总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAvailabilityZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZonesResponse", string(data)}, " ")
}
