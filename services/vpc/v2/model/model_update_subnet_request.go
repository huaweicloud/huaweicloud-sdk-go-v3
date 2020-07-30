/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type UpdateSubnetRequest struct {
	VpcId string `json:"vpc_id"`
	SubnetId string `json:"subnet_id"`
	Body *UpdateSubnetRequestBody `json:"body,omitempty"`
}
