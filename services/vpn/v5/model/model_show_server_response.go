package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowServerResponse struct {

	// 服务端 ID
	Id *string `json:"id,omitempty"`

	// P2C VPN 网关 ID
	P2cVgwId *string `json:"p2c_vgw_id,omitempty"`

	// 客户端网段
	ClientCidr *string `json:"client_cidr,omitempty"`

	// 本端网段列表
	LocalSubnets *[]string `json:"local_subnets,omitempty"`

	// 客户端认证类型
	ClientAuthType *string `json:"client_auth_type,omitempty"`

	// 隧道协议类型
	TunnelProtocol *string `json:"tunnel_protocol,omitempty"`

	ServerCertificate *ShowServerResponseServerCertificate `json:"server_certificate,omitempty"`

	ClientCaCertificates *[]QueryClientCaCertificateBody `json:"client_ca_certificates,omitempty"`

	SslOptions *ShowServerResponseSslOptions `json:"ssl_options,omitempty"`

	// 服务端状态
	Status *string `json:"status,omitempty"`

	// 创建时间
	CreatedAt *sdktime.SdkTime `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *sdktime.SdkTime `json:"updated_at,omitempty"`
}

func (o ShowServerResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerResponse struct{}"
	}

	return strings.Join([]string{"ShowServerResponse", string(data)}, " ")
}
