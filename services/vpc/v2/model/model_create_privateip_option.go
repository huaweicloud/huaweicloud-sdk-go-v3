/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type CreatePrivateipOption struct {
	// 分配IP的子网标识
	SubnetId string `json:"subnet_id"`
	// 功能说明：指定IP地址申请 取值范围：子网段中的可以使用且未分配的IP地址，不指定时由系统自动分配
	IpAddress string `json:"ip_address,omitempty"`
}
