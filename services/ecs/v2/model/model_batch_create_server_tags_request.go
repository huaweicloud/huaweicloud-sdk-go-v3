/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type BatchCreateServerTagsRequest struct {
	ServerId string `json:"server_id"`
	Body *BatchCreateServerTagsRequestBody `json:"body,omitempty"`
}
