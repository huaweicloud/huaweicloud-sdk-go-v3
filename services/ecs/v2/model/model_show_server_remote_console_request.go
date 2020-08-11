/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type ShowServerRemoteConsoleRequest struct {
	ServerId string `json:"server_id"`
	Body *ShowServerRemoteConsoleRequestBody `json:"body,omitempty"`
}
