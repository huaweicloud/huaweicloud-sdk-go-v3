package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAvailableZonesResponse struct {
	// 区域ID。

	RegionId *string `json:"region_id,omitempty"`
	// 可用分区数组。

	AvailableZones *[]AvailableZones `json:"available_zones,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}
