/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ResizePostPaidServerRequest struct {
	ServerId string `json:"server_id"`
	Body *ResizePostPaidServerRequestBody `json:"body,omitempty"`
}
