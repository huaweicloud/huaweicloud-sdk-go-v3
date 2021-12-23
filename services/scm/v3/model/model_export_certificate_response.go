package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportCertificateResponse struct {
	// 证书内容。

	Certificate *string `json:"certificate,omitempty"`
	// 证书链。

	CertificateChain *string `json:"certificate_chain,omitempty"`
	// 证书私钥。

	PrivateKey     *string `json:"private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateResponse", string(data)}, " ")
}
