package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowVpnGatewayRoutingTableResponse Response Object
type ShowVpnGatewayRoutingTableResponse struct {
	RoutingTable *[]VpnGatewayRoutingTableEntryVo `json:"routing_table,omitempty"`

	TotalCount *int64 `json:"total_count,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowVpnGatewayRoutingTableResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowVpnGatewayRoutingTableResponse struct{}"
	}

	return strings.Join([]string{"ShowVpnGatewayRoutingTableResponse", string(data)}, " ")
}
