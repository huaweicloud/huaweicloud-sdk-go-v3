/*
    * VPC
    *
    * VPC Open API
    *
*/

package model

// Response Object
type ListVpcPeeringsResponse struct {
	// peering对象列表
	Peerings []VpcPeering `json:"peerings,omitempty"`
	// 分页信息
	PeeringsLinks []NeutronPageLink `json:"peerings_links,omitempty"`
}
