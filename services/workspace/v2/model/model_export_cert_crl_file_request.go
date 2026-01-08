package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertCrlFileRequest Request Object
type ExportCertCrlFileRequest struct {

	// 证书id。
	CertId string `json:"cert_id"`
}

func (o ExportCertCrlFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertCrlFileRequest struct{}"
	}

	return strings.Join([]string{"ExportCertCrlFileRequest", string(data)}, " ")
}
