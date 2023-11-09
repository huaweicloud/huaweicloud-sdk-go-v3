package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateVgwResponse Response Object
type UpdateVgwResponse struct {
	VpnGateway *UpdateResponseVpnGateway `json:"vpn_gateway,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o UpdateVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVgwResponse struct{}"
	}

	return strings.Join([]string{"UpdateVgwResponse", string(data)}, " ")
}
