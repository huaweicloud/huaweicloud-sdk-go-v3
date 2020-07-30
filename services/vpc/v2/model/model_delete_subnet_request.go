/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type DeleteSubnetRequest struct {
	VpcId string `json:"vpc_id"`
	SubnetId string `json:"subnet_id"`
}
