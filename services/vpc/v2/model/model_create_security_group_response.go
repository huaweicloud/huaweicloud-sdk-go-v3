/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type CreateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}
