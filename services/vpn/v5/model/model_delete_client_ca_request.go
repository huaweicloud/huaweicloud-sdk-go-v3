package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteClientCaRequest Request Object
type DeleteClientCaRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 客户端 CA 证书 ID
	ClientCaCertificateId string `json:"client_ca_certificate_id"`
}

func (o DeleteClientCaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteClientCaRequest struct{}"
	}

	return strings.Join([]string{"DeleteClientCaRequest", string(data)}, " ")
}
