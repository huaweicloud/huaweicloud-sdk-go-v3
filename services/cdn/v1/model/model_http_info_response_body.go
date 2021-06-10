package model

import (
	"encoding/json"

	"strings"
)

type HttpInfoResponseBody struct {
	// HTTPS证书是否启用。0：不启用，此时无需填写证书及私钥参数；1：启用HTTPS加速并协议跟随回源；2：启用HTTPS加速并HTTP回源，开启时需要传递证书及私钥

	HttpsStatus *string `json:"https_status,omitempty"`
	// 证书名称。（长度限制为3-32字符）。

	CertName *string `json:"cert_name,omitempty"`
	// 证书内容。

	Certificate *string `json:"certificate,omitempty"`
	// 功能说明： HTTPS协议使用的私钥，不启用证书则无需输入。  取值范围：PEM编码格式。

	PrivateKey *string `json:"private_key,omitempty"`
	// 证书类型。1：代表华为云托管证书；0：表示自有证书。

	CertificateType *int32 `json:"certificate_type,omitempty"`
	// 客户端请求是否强制重定向。1是，0否。

	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	ForceRedirectConfig *ForceRedirect `json:"force_redirect_config,omitempty"`
	// 是否使用HTTP2.0。（1是，0否。）

	Http2 *int32 `json:"http2,omitempty"`
	// 证书过期时间

	ExpirationTime *int64 `json:"expiration_time,omitempty"`
}

func (o HttpInfoResponseBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "HttpInfoResponseBody struct{}"
	}

	return strings.Join([]string{"HttpInfoResponseBody", string(data)}, " ")
}
