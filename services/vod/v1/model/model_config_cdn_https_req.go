package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConfigCdnHttpsReq struct {

	// 来源，user表示用户自己上传，scm表示scm证书，默认user
	Source *string `json:"source,omitempty"`

	// 加速域名
	Domain string `json:"domain"`

	// 证书名称，若来源是scm则非必填，不填默认取scm上的证书名称
	CertName *string `json:"cert_name,omitempty"`

	// scm证书ID，若来源是scm则必填
	CertId *string `json:"cert_id,omitempty"`

	// https配置，0为关闭https配置，1为开启https配置，默认0
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// 证书内容，若来源是user则需填写，来源是scm则非必填
	Certificate *string `json:"certificate,omitempty"`

	// 私钥，若来源是user则需填写，来源是scm则非必填
	PrivateKey *string `json:"private_key,omitempty"`

	// 客户端请求是否强制重定向，0表示不重定向，1表示重定向，默认0
	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	// 是否使用HTTP2.0，0表示不使用HTTP2.0，1表示使用，默认0
	Http2 *int32 `json:"http2,omitempty"`
}

func (o ConfigCdnHttpsReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConfigCdnHttpsReq struct{}"
	}

	return strings.Join([]string{"ConfigCdnHttpsReq", string(data)}, " ")
}
