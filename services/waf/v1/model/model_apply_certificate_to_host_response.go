package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ApplyCertificateToHostResponse struct {

	// 证书id
	Id *string `json:"id,omitempty" xml:"id"`

	// 证书名
	Name *string `json:"name,omitempty" xml:"name"`

	// 时间戳
	Timestamp *int64 `json:"timestamp,omitempty" xml:"timestamp"`

	// 过期时间
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time"`

	// 绑定域名列表
	BindHost       *[]CertificateBundingHostBody `json:"bind_host,omitempty" xml:"bind_host"`
	HttpStatusCode int                           `json:"-"`
}

func (o ApplyCertificateToHostResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApplyCertificateToHostResponse struct{}"
	}

	return strings.Join([]string{"ApplyCertificateToHostResponse", string(data)}, " ")
}
