package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TlsConfig TLS信息。
type TlsConfig struct {

	// pem内容, 有则更新，无则上传。查询不返回。
	CertPem *string `json:"cert_pem,omitempty"`

	// 证书生效开始时间，时间参考样例 2022-01-25T09:24:27。
	CertStartTime *string `json:"cert_start_time,omitempty"`

	// 证书生效截止时间，时间参考样例 2022-01-25T09:24:27。
	CertEndTime *string `json:"cert_end_time,omitempty"`
}

func (o TlsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TlsConfig struct{}"
	}

	return strings.Join([]string{"TlsConfig", string(data)}, " ")
}
