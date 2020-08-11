/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaListServersDetailsRequest struct {
	ChangesSince string `json:"changes-since,omitempty"`
	Flavor string `json:"flavor,omitempty"`
	Image string `json:"image,omitempty"`
	Ip string `json:"ip,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	Name string `json:"name,omitempty"`
	NotTags string `json:"not-tags,omitempty"`
	ReservationId string `json:"reservation_id,omitempty"`
	SortKey string `json:"sort_key,omitempty"`
	Status string `json:"status,omitempty"`
	Tags string `json:"tags,omitempty"`
	OpenStackAPIVersion string `json:"OpenStack-API-Version,omitempty"`
}
