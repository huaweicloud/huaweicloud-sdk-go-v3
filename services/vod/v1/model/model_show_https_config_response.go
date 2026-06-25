package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpsConfigResponse Response Object
type ShowHttpsConfigResponse struct {

	// 来源，user表示用户自己上传，scm表示scm证书
	Source *string `json:"source,omitempty"`

	// 证书名称
	CertName *string `json:"cert_name,omitempty"`

	// 证书id
	CertId *string `json:"cert_id,omitempty"`

	// https配置
	HttpsStatus *int32 `json:"https_status,omitempty"`

	// 证书内容
	Certificate *string `json:"certificate,omitempty"`

	// 客户端请求是否强制重定向，0表示不重定向，1表示重定向
	ForceRedirectHttps *int32 `json:"force_redirect_https,omitempty"`

	// 是否使用HTTP2.0，0表示不使用HTTP2.0，1表示使用
	Http2          *int32 `json:"http2,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowHttpsConfigResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpsConfigResponse struct{}"
	}

	return strings.Join([]string{"ShowHttpsConfigResponse", string(data)}, " ")
}
