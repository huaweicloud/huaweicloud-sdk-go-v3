/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ResizeServerRequest struct {
	ServerId string `json:"server_id"`
	Body *ResizeServerRequestBody `json:"body,omitempty"`
}
