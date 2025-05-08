package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAvailableZonesResponse Response Object
type ListAvailableZonesResponse struct {

	// **参数解释**： 区域ID。 **取值范围**： 不涉及。
	RegionId *string `json:"region_id,omitempty"`

	// **参数解释**： 可用区数组。
	AvailableZones *[]AvailableZonesResp `json:"available_zones,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAvailableZonesResponse struct{}"
	}

	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}
