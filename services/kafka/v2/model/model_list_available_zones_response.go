/*
 * Kafka
 *
 * Kafka Document API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListAvailableZonesResponse struct {
	// 区域ID。
	RegionId *string `json:"region_id,omitempty"`
	// 可用区数组。
	AvailableZones *[]ListAvailableZonesRespAvailableZones `json:"available_zones,omitempty"`
	HttpStatusCode int                                     `json:"-"`
}

func (o ListAvailableZonesResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListAvailableZonesResponse", string(data)}, " ")
}
