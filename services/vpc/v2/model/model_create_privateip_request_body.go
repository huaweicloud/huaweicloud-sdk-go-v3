/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type CreatePrivateipRequestBody struct {
	// 私有IP列表对象
	Privateips []CreatePrivateipOption `json:"privateips"`
}
