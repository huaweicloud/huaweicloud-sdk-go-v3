/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Request Object
type ListPrivateipsRequest struct {
	SubnetId string `json:"subnet_id"`
	Limit int32 `json:"limit,omitempty"`
	Marker string `json:"marker,omitempty"`
}
