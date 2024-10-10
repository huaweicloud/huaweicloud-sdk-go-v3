package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainCertificateResponse Response Object
type ShowDomainCertificateResponse struct {

	// 域名id
	DomainId *string `json:"domain_id,omitempty"`

	// 域名
	DomainName *string `json:"domain_name,omitempty"`

	CertInfo       *CertInfo `json:"cert_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowDomainCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainCertificateResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainCertificateResponse", string(data)}, " ")
}
