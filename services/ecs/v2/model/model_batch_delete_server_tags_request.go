/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type BatchDeleteServerTagsRequest struct {
	ServerId string `json:"server_id"`
	Body *BatchDeleteServerTagsRequestBody `json:"body,omitempty"`
}
