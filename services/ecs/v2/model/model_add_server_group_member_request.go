/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type AddServerGroupMemberRequest struct {
	ServerGroupId string `json:"server_group_id"`
	Body *AddServerGroupMemberRequestBody `json:"body,omitempty"`
}
