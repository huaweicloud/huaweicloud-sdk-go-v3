/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// 可用域信息
type NovaAvailabilityZone struct {
	// 该字段的值为null。
	Hosts []string `json:"hosts"`
	// 可用域的名称。
	ZoneName string `json:"zoneName"`
	ZoneState *NovaAvailabilityZoneState `json:"zoneState"`
}
