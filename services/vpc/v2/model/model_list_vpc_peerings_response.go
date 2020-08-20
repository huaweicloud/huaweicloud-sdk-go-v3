/*
 * VPC
 *
 * VPC Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListVpcPeeringsResponse struct {
	// peering对象列表
	Peerings []VpcPeering `json:"peerings,omitempty"`
	// 分页信息
	PeeringsLinks []NeutronPageLink `json:"peerings_links,omitempty"`
}

func (o ListVpcPeeringsResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListVpcPeeringsResponse", string(data)}, " ")
}
