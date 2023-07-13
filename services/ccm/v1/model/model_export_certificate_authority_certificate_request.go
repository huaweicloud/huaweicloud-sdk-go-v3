package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertificateAuthorityCertificateRequest Request Object
type ExportCertificateAuthorityCertificateRequest struct {

	// 所要导出的CA证书ID。
	CaId string `json:"ca_id"`
}

func (o ExportCertificateAuthorityCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateAuthorityCertificateRequest struct{}"
	}

	return strings.Join([]string{"ExportCertificateAuthorityCertificateRequest", string(data)}, " ")
}
