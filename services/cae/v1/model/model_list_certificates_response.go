package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCertificatesResponse Response Object
type ListCertificatesResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *CertificateKindObj `json:"kind,omitempty"`

	// 证书列表。
	Items          *[]CertItem `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ListCertificatesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCertificatesResponse struct{}"
	}

	return strings.Join([]string{"ListCertificatesResponse", string(data)}, " ")
}
