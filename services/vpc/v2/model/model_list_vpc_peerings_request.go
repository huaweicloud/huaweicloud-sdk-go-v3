/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListVpcPeeringsRequest struct {
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
	VpcId string `json:"vpc_id,omitempty"`
}
