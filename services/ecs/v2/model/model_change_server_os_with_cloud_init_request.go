/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ChangeServerOsWithCloudInitRequest struct {
	ServerId string `json:"server_id"`
	Body *ChangeServerOsWithCloudInitRequestBody `json:"body,omitempty"`
}
