/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaListKeypairsRequest struct {
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	OpenStackAPIVersion string `json:"OpenStack-API-Version,omitempty"`
}
