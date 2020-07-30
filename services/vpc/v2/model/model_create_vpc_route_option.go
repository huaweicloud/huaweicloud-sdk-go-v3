/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// 
type CreateVpcRouteOption struct {
	// 路由目的地址CIDR，如192.168.200.0/24。
	Destination string `json:"destination"`
	// 功能说明：路由下一跳  取值范围：如果type为peering类型，则nexthop为peering的ID
	Nexthop string `json:"nexthop"`
	// 功能说明：路由类型  取值范围：peering
	Type string `json:"type"`
	// 请求添加路由的VPC ID
	VpcId string `json:"vpc_id"`
}
