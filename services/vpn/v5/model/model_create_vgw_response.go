package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateVgwResponse Response Object
type CreateVgwResponse struct {
	VpnGateway *CreateResponseVpnGateway `json:"vpn_gateway,omitempty"`

	// 请求ID
	RequestId *string `json:"request_id,omitempty"`

	HeaderResponseToken *string `json:"header-response-token,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o CreateVgwResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVgwResponse struct{}"
	}

	return strings.Join([]string{"CreateVgwResponse", string(data)}, " ")
}
