/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaAssociateSecurityGroupRequest struct {
	ServerId string `json:"server_id"`
	Body *NovaAssociateSecurityGroupRequestBody `json:"body,omitempty"`
}
