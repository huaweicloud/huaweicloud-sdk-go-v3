/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ResetServerPasswordRequest struct {
	ServerId string `json:"server_id"`
	Body *ResetServerPasswordRequestBody `json:"body,omitempty"`
}
