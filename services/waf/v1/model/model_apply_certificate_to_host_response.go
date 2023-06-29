package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ApplyCertificateToHostResponse Response Object
type ApplyCertificateToHostResponse struct {

	// 证书id
	Id *string `json:"id,omitempty"`

	// 证书名
	Name *string `json:"name,omitempty"`

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 过期时间
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 绑定域名列表
	BindHost       *[]CertificateBundingHostBody `json:"bind_host,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ApplyCertificateToHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostResponse struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostResponse", string(data)}, " ")
}
