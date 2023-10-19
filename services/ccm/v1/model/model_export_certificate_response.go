package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertificateResponse Response Object
type ExportCertificateResponse struct {

	// 私钥内容。
	PrivateKey *string `json:"private_key,omitempty"`

	// 证书内容。
	Certificate *string `json:"certificate,omitempty"`

	// 证书链内容。
	CertificateChain *string `json:"certificate_chain,omitempty"`

	// 加密证书内容。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// 加密私钥内容。
	EncPrivateKey *string `json:"enc_private_key,omitempty"`

	// 加密私钥的国密GMT0009标准规范SM2数字信封。
	EncSm2EnvelopedKey *string `json:"enc_sm2_enveloped_key,omitempty"`

	// 加密私钥的国密GMT0010标准规范签名数字信封。
	SignedAndEnvelopedData *string `json:"signed_and_enveloped_data,omitempty"`
	HttpStatusCode         int     `json:"-"`
}

func (o ExportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateResponse", string(data)}, " ")
}
