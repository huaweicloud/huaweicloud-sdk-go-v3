/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type UpdateVpcRequest struct {
	VpcId string `json:"vpc_id"`
	Body *UpdateVpcRequestBody `json:"body,omitempty"`
}
