package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ServerCertificateConfig 服务端证书配置结构体
type ServerCertificateConfig struct {

	// 是否开启服务端OCSP装订
	OcspStaplingEnable *bool `json:"ocsp_stapling_enable,omitempty"`

	// ocsp服务器端CA证书id，仅当ocsp服务器为https协议时支持配置。
	OcspServerCaId *string `json:"ocsp_server_ca_id,omitempty"`

	// 访问ocsp服务器是否开启SSL
	OcspSslEnable *bool `json:"ocsp_ssl_enable,omitempty"`
}

func (o ServerCertificateConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ServerCertificateConfig struct{}"
	}

	return strings.Join([]string{"ServerCertificateConfig", string(data)}, " ")
}
