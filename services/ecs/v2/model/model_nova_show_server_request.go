/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaShowServerRequest struct {
	ServerId string `json:"server_id"`
	OpenStackAPIVersion string `json:"OpenStack-API-Version,omitempty"`
}
