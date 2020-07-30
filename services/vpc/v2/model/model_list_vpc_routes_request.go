/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListVpcRoutesRequest struct {
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	Id string `json:"id,omitempty"`
	Type string `json:"type,omitempty"`
	VpcId string `json:"vpc_id,omitempty"`
	Destination string `json:"destination,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
}
