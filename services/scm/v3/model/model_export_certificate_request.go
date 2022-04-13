package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportCertificateRequest struct {
	// 证书id。

	CertificateId string `json:"certificate_id"`
}

func (o ExportCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateRequest struct{}"
	}

	return strings.Join([]string{"ExportCertificateRequest", string(data)}, " ")
}
