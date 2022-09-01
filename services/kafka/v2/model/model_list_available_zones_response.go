package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAvailableZonesResponse struct {

	// 区域ID。
	RegionId *string `json:"region_id,omitempty" xml:"region_id"`

	// 可用区数组。
	AvailableZones *[]ListAvailableZonesRespAvailableZones `json:"available_zones,omitempty" xml:"available_zones"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}
