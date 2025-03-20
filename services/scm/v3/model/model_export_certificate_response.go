package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertificateResponse Response Object
type ExportCertificateResponse struct {

	// 证书及证书链。
	EntireCertificate *string `json:"entire_certificate,omitempty"`

	// 证书内容，不包含证书链。
	Certificate *string `json:"certificate,omitempty"`

	// 证书链。
	CertificateChain *string `json:"certificate_chain,omitempty"`

	// 证书私钥。
	PrivateKey *string `json:"private_key,omitempty"`

	// 国密证书返回，加密证书内容。
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// 国密证书返回，加密证书私钥。 - 自己生成csr的方式：该字段是加密状态，需要解密。   解密思路如下:    1、从数字信封中获取“对称密钥密文”和“加密私钥密文”；   2、使用签名私钥解密“对称密钥密文”，得到“对称密钥明文”；   3、使用对称密钥解密“加密私钥密文”，得到“加密私钥明文”。 - 非自己生成csr的方式：该字段不需要解密。
	EncPrivateKey  *string `json:"enc_private_key,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ExportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ExportCertificateResponse", string(data)}, " ")
}
