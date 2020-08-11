/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ReinstallServerWithCloudInitRequest struct {
	ServerId string `json:"server_id"`
	Body *ReinstallServerWithCloudInitRequestBody `json:"body,omitempty"`
}
