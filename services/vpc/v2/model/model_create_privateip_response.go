/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type CreatePrivateipResponse struct {
	// 私有IP列表对象
	Privateips []Privateip `json:"privateips,omitempty"`
}
