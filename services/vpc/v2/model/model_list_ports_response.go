/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListPortsResponse struct {
	// port列表对象
	Ports []Port `json:"ports,omitempty"`
}
