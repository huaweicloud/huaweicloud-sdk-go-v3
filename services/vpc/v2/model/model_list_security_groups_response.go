/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListSecurityGroupsResponse struct {
	// 安全组列表对象
	SecurityGroups []SecurityGroup `json:"security_groups,omitempty"`
}
