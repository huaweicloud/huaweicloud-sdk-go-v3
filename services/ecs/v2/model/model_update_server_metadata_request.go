/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type UpdateServerMetadataRequest struct {
	ServerId string `json:"server_id"`
	Body *UpdateServerMetadataRequestBody `json:"body,omitempty"`
}
