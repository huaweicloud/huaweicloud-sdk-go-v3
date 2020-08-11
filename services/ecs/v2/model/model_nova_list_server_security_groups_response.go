/*
    * ecs
    *
    * ECS Open API
    *
*/

package model

// Response Object
type NovaListServerSecurityGroupsResponse struct {
	// security_group列表
	SecurityGroups []NovaSecurityGroup `json:"security_groups,omitempty"`
}
