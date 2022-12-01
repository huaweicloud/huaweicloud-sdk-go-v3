package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 数据库SSL证书信息体。如果数据库启用了SSL安全连接，请确保相关配置正确，并输入SSL证书相关参数，否则无需填写此参数。
type EndpointSslConfig struct {

	// 是否SSL安全连接。如果数据库启用了SSL安全连接，参数值为true。
	SslLink bool `json:"ssl_link"`

	// SSL证书名字。
	SslCertName string `json:"ssl_cert_name"`

	// SSL证书内容，用base64加密。
	SslCertKey string `json:"ssl_cert_key"`

	// SSL证书内容checksum值，后端校验，源库安全连接必选。
	SslCertCheckSum string `json:"ssl_cert_check_sum"`

	// SSL证书密码，证书文件后缀为.p12时必填。
	SslCertPassword *string `json:"ssl_cert_password,omitempty"`
}

func (o EndpointSslConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EndpointSslConfig struct{}"
	}

	return strings.Join([]string{"EndpointSslConfig", string(data)}, " ")
}
