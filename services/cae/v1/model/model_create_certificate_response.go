package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateResponse Response Object
type CreateCertificateResponse struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion *string `json:"api_version,omitempty"`

	// API类型，固定值“Certificate”，该值不可修改。
	Kind *string `json:"kind,omitempty"`

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
