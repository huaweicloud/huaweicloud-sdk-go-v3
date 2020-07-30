/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ShowVpcPeeringResponse struct {
	Peering *VpcPeering `json:"peering,omitempty"`
}
