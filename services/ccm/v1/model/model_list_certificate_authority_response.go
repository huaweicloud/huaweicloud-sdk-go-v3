package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificateAuthorityResponse Response Object
type ListCertificateAuthorityResponse struct {

	// CA证书总数。
	Total *int32 `json:"total,omitempty"`

	// CA列表，详情请参见**CertificateAuthorities**字段数据结构说明。
	CertificateAuthorities *[]CertificateAuthorities `json:"certificate_authorities,omitempty"`
	HttpStatusCode         int                       `json:"-"`
}

func (o ListCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"ListCertificateAuthorityResponse", string(data)}, " ")
}
