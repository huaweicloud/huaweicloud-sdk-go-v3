package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListCertificatesResponse struct {

	// 证书列表
	Items *[]CertificateBody `json:"items,omitempty"`

	// 证书总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}
