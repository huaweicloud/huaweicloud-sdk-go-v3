package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePeerLinkRequest Request Object
type DeletePeerLinkRequest struct {

	// 全域接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 全域接入网关对等体
	PeerLinkId string `json:"peer_link_id"`
}

func (o DeletePeerLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePeerLinkRequest struct{}"
	}

	return strings.Join([]string{"DeletePeerLinkRequest", string(data)}, " ")
}
