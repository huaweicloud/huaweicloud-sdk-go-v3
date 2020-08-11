/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type UpdateServerRequest struct {
	ServerId string `json:"server_id"`
	Body *UpdateServerRequestBody `json:"body,omitempty"`
}
