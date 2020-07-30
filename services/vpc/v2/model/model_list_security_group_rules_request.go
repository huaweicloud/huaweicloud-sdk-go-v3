/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListSecurityGroupRulesRequest struct {
	Marker string `json:"marker,omitempty"`
	Limit int32 `json:"limit,omitempty"`
	SecurityGroupId string `json:"security_group_id,omitempty"`
}
