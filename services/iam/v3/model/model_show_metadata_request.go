package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowMetadataRequest struct {
	// 身份提供商ID。

	IdpId string `json:"idp_id"`
	// 协议ID。

	ProtocolId string `json:"protocol_id"`
}

func (o ShowMetadataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMetadataRequest struct{}"
	}

	return strings.Join([]string{"ShowMetadataRequest", string(data)}, " ")
}
