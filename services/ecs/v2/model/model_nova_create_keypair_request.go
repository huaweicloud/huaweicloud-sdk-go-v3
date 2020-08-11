/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type NovaCreateKeypairRequest struct {
	OpenStackAPIVersion string `json:"OpenStack-API-Version,omitempty"`
	Body *NovaCreateKeypairRequestBody `json:"body,omitempty"`
}
