package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcPeeringsResponse Response Object
type ListVpcPeeringsResponse struct {

	// peering对象列表
	Peerings *[]VpcPeering `json:"peerings,omitempty"`

	// 分页信息
	PeeringsLinks  *[]NeutronPageLink `json:"peerings_links,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVpcPeeringsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcPeeringsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcPeeringsResponse", string(data)}, " ")
}
