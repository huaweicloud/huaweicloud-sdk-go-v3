package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceSslDetailResponse Response Object
type ShowInstanceSslDetailResponse struct {

	// 开启或关闭SSL。true：开启/false：关闭
	Enabled *bool `json:"enabled,omitempty"`

	// SSL连接IP。
	Ip *string `json:"ip,omitempty"`

	// SSL连接端口。
	Port *string `json:"port,omitempty"`

	// SSL连接域名。
	DomainName *string `json:"domain_name,omitempty"`

	// SSL证书有效期（UTC时间）。
	SslExpiredAt *string `json:"ssl_expired_at,omitempty"`

	// SSL证书是否有效。
	SslValidated   *bool `json:"ssl_validated,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowInstanceSslDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceSslDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceSslDetailResponse", string(data)}, " ")
}
