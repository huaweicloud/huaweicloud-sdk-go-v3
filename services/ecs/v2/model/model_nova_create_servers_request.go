/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaCreateServersRequest struct {
	OpenStackAPIVersion string `json:"OpenStack-API-Version,omitempty"`
	Body *NovaCreateServersRequestBody `json:"body,omitempty"`
}
