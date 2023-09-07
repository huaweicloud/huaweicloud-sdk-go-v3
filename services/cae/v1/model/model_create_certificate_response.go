package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateResponse Response Object
type CreateCertificateResponse struct {
	ApiVersion *ApiVersionObj `json:"api_version,omitempty"`

	Kind *CertificateKindObj `json:"kind,omitempty"`

	// 证书列表。
	Items          *[]CertItem `json:"items,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateResponse", string(data)}, " ")
}
