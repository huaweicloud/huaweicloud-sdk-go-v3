package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ExportCertificateResponse struct {

	// 证书内容。
	Certificate *string `json:"certificate,omitempty" xml:"certificate"`

	// 证书链。
	CertificateChain *string `json:"certificate_chain,omitempty" xml:"certificate_chain"`

	// 证书私钥。
	PrivateKey *string `json:"private_key,omitempty" xml:"private_key"`

	// 国密证书返回，加密证书内容。
	EncCertificate *string `json:"enc_certificate,omitempty" xml:"enc_certificate"`

	// 国密证书返回，加密证书私钥。
	EncPrivateKey  *string `json:"enc_private_key,omitempty" xml:"enc_private_key"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateResponse", string(data)}, " ")
}
