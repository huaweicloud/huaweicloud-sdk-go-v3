package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailabilityZoneResponse Response Object
type ListAvailabilityZoneResponse struct {

	// 云应用支持的可用分区列表。
	AvailabilityZones *[]AvailabilityZoneInfo `json:"availability_zones,omitempty"`

	// 总数
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListAvailabilityZoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailabilityZoneResponse struct{}"
	}

	return strings.Join([]string{"ListAvailabilityZoneResponse", string(data)}, " ")
}
