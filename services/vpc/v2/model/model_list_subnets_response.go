/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListSubnetsResponse struct {
	// subnet对象列表
	Subnets []Subnet `json:"subnets,omitempty"`
}
