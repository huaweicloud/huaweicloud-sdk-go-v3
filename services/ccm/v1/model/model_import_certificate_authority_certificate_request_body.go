package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportCertificateAuthorityCertificateRequestBody struct {

	// 证书内容。
	Certificate string `json:"certificate"`

	// 证书链内容。
	CertificateChain *string `json:"certificate_chain,omitempty"`
}

func (o ImportCertificateAuthorityCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificateAuthorityCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ImportCertificateAuthorityCertificateRequestBody", string(data)}, " ")
}
