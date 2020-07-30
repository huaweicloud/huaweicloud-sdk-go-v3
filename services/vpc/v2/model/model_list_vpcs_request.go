/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListVpcsRequest struct {
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	Id string `json:"id,omitempty"`
	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}
