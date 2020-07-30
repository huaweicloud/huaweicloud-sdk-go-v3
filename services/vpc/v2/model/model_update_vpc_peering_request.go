/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type UpdateVpcPeeringRequest struct {
	PeeringId string `json:"peering_id"`
	Body *UpdateVpcPeeringRequestBody `json:"body,omitempty"`
}
