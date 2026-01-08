package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertPemFileRequest Request Object
type ExportCertPemFileRequest struct {

	// 证书id。
	CertId string `json:"cert_id"`
}

func (o ExportCertPemFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertPemFileRequest struct{}"
	}

	return strings.Join([]string{"ExportCertPemFileRequest", string(data)}, " ")
}
