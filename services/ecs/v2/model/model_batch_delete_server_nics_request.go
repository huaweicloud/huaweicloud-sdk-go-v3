/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type BatchDeleteServerNicsRequest struct {
	ServerId string `json:"server_id"`
	Body *BatchDeleteServerNicsRequestBody `json:"body,omitempty"`
}
