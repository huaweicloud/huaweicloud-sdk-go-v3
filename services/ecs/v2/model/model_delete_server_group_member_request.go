/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Request Object
type DeleteServerGroupMemberRequest struct {
	ServerGroupId string `json:"server_group_id"`
	Body *DeleteServerGroupMemberRequestBody `json:"body,omitempty"`
}
