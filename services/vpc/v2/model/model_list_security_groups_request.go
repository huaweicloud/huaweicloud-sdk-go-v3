/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListSecurityGroupsRequest struct {
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	VpcId string `json:"vpc_id,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}
