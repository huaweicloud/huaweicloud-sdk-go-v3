package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ImportCertificateAuthorityCertificateRequest struct {

	// 所要导入的CA证书ID。
	CaId string `json:"ca_id"`

	Body *ImportCertificateAuthorityCertificateRequestBody `json:"body,omitempty"`
}

func (o ImportCertificateAuthorityCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificateAuthorityCertificateRequest struct{}"
	}

	return strings.Join([]string{"ImportCertificateAuthorityCertificateRequest", string(data)}, " ")
}
