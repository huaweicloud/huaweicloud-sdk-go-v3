package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpcPeeringInfosResponse Response Object
type ListVpcPeeringInfosResponse struct {

	// 对等连接信息对象列表
	PeeringInfos *[]VpcPeering `json:"peering_infos,omitempty"`

	// 分页信息
	PeeringInfosLinks *[]NeutronPageLink `json:"peering_infos_links,omitempty"`
	HttpStatusCode    int                `json:"-"`
}

func (o ListVpcPeeringInfosResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpcPeeringInfosResponse struct{}"
	}

	return strings.Join([]string{"ListVpcPeeringInfosResponse", string(data)}, " ")
}
