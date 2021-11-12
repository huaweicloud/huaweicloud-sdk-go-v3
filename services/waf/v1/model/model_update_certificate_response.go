package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type UpdateCertificateResponse struct {
	// 证书ID

	Id *string `json:"id,omitempty"`
	// 证书名

	Name *string `json:"name,omitempty"`
	// 证书过期时间戳

	ExpireTime *int64 `json:"expire_time,omitempty"`
	// 时间戳

	Timestamp      *int64 `json:"timestamp,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o UpdateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateCertificateResponse", string(data)}, " ")
}
