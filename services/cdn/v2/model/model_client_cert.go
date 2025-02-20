package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClientCert 客户端证书配置
type ClientCert struct {

	// 客户端证书配置开关，on：打开；off：关闭。
	Status string `json:"status"`

	// 客户端CA证书的内容，仅支持PEM格式。
	TrustedCert string `json:"trusted_cert"`

	// 客户端CA证书指定的域名。   > 1. 如果不配置域名，则CDN会放行所有持有该CA证书的客户端请求。   > 2. 最多可配置100个域名，多个域名用“,”或“|”分隔。
	Hosts *string `json:"hosts,omitempty"`
}

func (o ClientCert) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClientCert struct{}"
	}

	return strings.Join([]string{"ClientCert", string(data)}, " ")
}
