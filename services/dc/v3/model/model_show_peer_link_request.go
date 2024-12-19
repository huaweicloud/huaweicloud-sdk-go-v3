package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPeerLinkRequest Request Object
type ShowPeerLinkRequest struct {

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`

	// 全域接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 全域接入网关对等体
	PeerLinkId string `json:"peer_link_id"`
}

func (o ShowPeerLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPeerLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowPeerLinkRequest", string(data)}, " ")
}
