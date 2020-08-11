/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaDisassociateSecurityGroupRequest struct {
	ServerId string `json:"server_id"`
	Body *NovaDisassociateSecurityGroupRequestBody `json:"body,omitempty"`
}
