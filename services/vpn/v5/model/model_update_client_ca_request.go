package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateClientCaRequest Request Object
type UpdateClientCaRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 客户端 CA 证书 ID
	ClientCaCertificateId string `json:"client_ca_certificate_id"`

	Body *UpdateClientCaCertificateRequestBody `json:"body,omitempty"`
}

func (o UpdateClientCaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientCaRequest struct{}"
	}

	return strings.Join([]string{"UpdateClientCaRequest", string(data)}, " ")
}
