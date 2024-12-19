package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePeerLinkRequest Request Object
type UpdatePeerLinkRequest struct {

	// 全域接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 全域接入网关对等体
	PeerLinkId string `json:"peer_link_id"`

	Body *UpdateExternalPeerLinkRequestBody `json:"body,omitempty"`
}

func (o UpdatePeerLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePeerLinkRequest struct{}"
	}

	return strings.Join([]string{"UpdatePeerLinkRequest", string(data)}, " ")
}
