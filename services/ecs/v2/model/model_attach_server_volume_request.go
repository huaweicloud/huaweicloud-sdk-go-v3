/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type AttachServerVolumeRequest struct {
	ServerId string `json:"server_id"`
	Body *AttachServerVolumeRequestBody `json:"body,omitempty"`
}
