package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClientCaRequest Request Object
type ShowClientCaRequest struct {

	// VPN服务端 ID
	VpnServerId string `json:"vpn_server_id"`

	// 客户端 CA 证书 ID
	ClientCaCertificateId string `json:"client_ca_certificate_id"`
}

func (o ShowClientCaRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClientCaRequest struct{}"
	}

	return strings.Join([]string{"ShowClientCaRequest", string(data)}, " ")
}
