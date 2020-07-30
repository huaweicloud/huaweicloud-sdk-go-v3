/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type UpdatePortRequest struct {
	PortId string `json:"port_id"`
	Body *UpdatePortRequestBody `json:"body,omitempty"`
}
