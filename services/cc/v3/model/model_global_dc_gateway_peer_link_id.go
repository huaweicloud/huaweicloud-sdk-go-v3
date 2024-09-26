package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GlobalDcGatewayPeerLinkId GDGW的连接ID。
type GlobalDcGatewayPeerLinkId struct {

	// GDGW的连接ID。
	GlobalDcGatewayPeerLinkId *string `json:"global_dc_gateway_peer_link_id,omitempty"`
}

func (o GlobalDcGatewayPeerLinkId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GlobalDcGatewayPeerLinkId struct{}"
	}

	return strings.Join([]string{"GlobalDcGatewayPeerLinkId", string(data)}, " ")
}
