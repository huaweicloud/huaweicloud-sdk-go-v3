package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListVpcPeeringsResponse struct {

	// peering对象列表
	Peerings *[]VpcPeering `json:"peerings,omitempty" xml:"peerings"`

	// 分页信息
	PeeringsLinks  *[]NeutronPageLink `json:"peerings_links,omitempty" xml:"peerings_links"`
	HttpStatusCode int                `json:"-"`
}

func (o ListVpcPeeringsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcPeeringsResponse struct{}"
	}

	return strings.Join([]string{"ListVpcPeeringsResponse", string(data)}, " ")
}
