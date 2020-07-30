/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListVpcsResponse struct {
	// vpc对象列表
	Vpcs []Vpc `json:"vpcs,omitempty"`
}
