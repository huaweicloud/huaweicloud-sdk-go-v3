package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateVpnGatewayCertificateRequestBodyContent struct {

	// VPN网关证书名称
	Name *string `json:"name,omitempty"`

	// 证书内容，国密证书时为签名证书内容
	Certificate *string `json:"certificate,omitempty"`

	// 证书私钥，国密证书时为签名证书私钥
	PrivateKey *string `json:"private_key,omitempty"`

	// VPN网关CA证书内容
	CertificateChain *string `json:"certificate_chain,omitempty"`

	// 国密证书的加密证书内容
	EncCertificate *string `json:"enc_certificate,omitempty"`

	// 国密证书的加密证书私钥
	EncPrivateKey *string `json:"enc_private_key,omitempty"`
}

func (o UpdateVpnGatewayCertificateRequestBodyContent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateVpnGatewayCertificateRequestBodyContent struct{}"
	}

	return strings.Join([]string{"UpdateVpnGatewayCertificateRequestBodyContent", string(data)}, " ")
}
