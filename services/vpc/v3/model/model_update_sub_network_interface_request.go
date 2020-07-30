/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type UpdateSubNetworkInterfaceRequest struct {
	SubNetworkInterfaceId string `json:"sub_network_interface_id"`
	Body *UpdateSubNetworkInterfaceRequestBody `json:"body,omitempty"`
}
