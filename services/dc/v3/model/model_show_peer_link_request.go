package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPeerLinkRequest Request Object
type ShowPeerLinkRequest struct {

	// 全球接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	// 全球接入网关对等体
	PeerLinkId string `json:"peer_link_id"`

	// 显示字段列表
	Fields *[]string `json:"fields,omitempty"`

	// show response ext-fields
	ExtFields *[]string `json:"ext_fields,omitempty"`
}

func (o ShowPeerLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPeerLinkRequest struct{}"
	}

	return strings.Join([]string{"ShowPeerLinkRequest", string(data)}, " ")
}
