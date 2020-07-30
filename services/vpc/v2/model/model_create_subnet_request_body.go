/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 创建子网对象
type CreateSubnetRequestBody struct {
	Subnet *CreateSubnetOption `json:"subnet"`
}
