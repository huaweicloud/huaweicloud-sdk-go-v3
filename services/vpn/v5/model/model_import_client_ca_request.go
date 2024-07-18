package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportClientCaRequest Request Object
type ImportClientCaRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 幂等性标识
	XClientToken *string `json:"X-Client-Token,omitempty"`

	Body *ImportClientCaCertificateRequestBody `json:"body,omitempty"`
}

func (o ImportClientCaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportClientCaRequest struct{}"
	}

	return strings.Join([]string{"ImportClientCaRequest", string(data)}, " ")
}
