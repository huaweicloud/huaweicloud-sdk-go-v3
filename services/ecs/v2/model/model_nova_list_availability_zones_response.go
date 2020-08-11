/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type NovaListAvailabilityZonesResponse struct {
	// 可用域信息。
	AvailabilityZoneInfo []NovaAvailabilityZone `json:"availabilityZoneInfo,omitempty"`
}
