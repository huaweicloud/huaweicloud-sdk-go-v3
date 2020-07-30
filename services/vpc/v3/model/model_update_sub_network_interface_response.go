/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type UpdateSubNetworkInterfaceResponse struct {
	// 请求ID
	RequestId string `json:"request_id,omitempty"`
	SubNetworkInterface *SubNetworkInterface `json:"sub_network_interface,omitempty"`
}
