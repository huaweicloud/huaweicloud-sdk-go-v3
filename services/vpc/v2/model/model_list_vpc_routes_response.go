/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListVpcRoutesResponse struct {
	// route对象列表
	Routes []VpcRoute `json:"routes,omitempty"`
	// 分页信息
	RoutesLinks []NeutronPageLink `json:"routes_links,omitempty"`
}
