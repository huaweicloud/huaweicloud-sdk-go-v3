/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type BatchAddServerNicsRequest struct {
	ServerId string `json:"server_id"`
	Body *BatchAddServerNicsRequestBody `json:"body,omitempty"`
}
