/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type CreateVpcPeeringResponse struct {
	Peering *VpcPeering `json:"peering,omitempty"`
}
