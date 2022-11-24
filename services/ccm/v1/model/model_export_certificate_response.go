package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportCertificateResponse struct {

	// 私钥内容。
	PrivateKey *string `json:"private_key,omitempty"`

	// 证书内容。
	Certificate *string `json:"certificate,omitempty"`

	// 证书链内容。
	CertificateChain *string `json:"certificate_chain,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o ExportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateResponse", string(data)}, " ")
}
