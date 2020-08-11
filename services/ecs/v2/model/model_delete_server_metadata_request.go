/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type DeleteServerMetadataRequest struct {
	Key string `json:"key"`
	ServerId string `json:"server_id"`
}
