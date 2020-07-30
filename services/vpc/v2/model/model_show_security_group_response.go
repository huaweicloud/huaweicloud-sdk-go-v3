/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ShowSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}
