package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertificateResponse Response Object
type DeleteCertificateResponse struct {

	// 证书ID
	Id *string `json:"id,omitempty"`

	// 证书名
	Name *string `json:"name,omitempty"`

	// 证书过期时间戳
	ExpireTime *int64 `json:"expire_time,omitempty"`

	// 证书上传时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 证书类型
	CertType       *string `json:"cert_type,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertificateResponse", string(data)}, " ")
}
