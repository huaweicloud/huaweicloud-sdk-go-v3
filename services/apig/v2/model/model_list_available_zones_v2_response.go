package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListAvailableZonesV2Response struct {

	// 可用区列表
	AvailableZones *[]AvailableZone `json:"available_zones,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListAvailableZonesV2Response) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesV2Response struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesV2Response", string(data)}, " ")
}
