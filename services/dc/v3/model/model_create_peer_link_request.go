package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePeerLinkRequest Request Object
type CreatePeerLinkRequest struct {

	// 全域接入网关ID
	GlobalDcGatewayId string `json:"global_dc_gateway_id"`

	Body *CreateExternalPeerLinkRequestBody `json:"body,omitempty"`
}

func (o CreatePeerLinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePeerLinkRequest struct{}"
	}

	return strings.Join([]string{"CreatePeerLinkRequest", string(data)}, " ")
}
