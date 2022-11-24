package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExportCertificateRequest struct {

	// 所要导出的私有证书ID。
	CertificateId string `json:"certificate_id"`

	Body *ExportCertificateRequestBody `json:"body,omitempty"`
}

func (o ExportCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertificateRequest struct{}"
	}

	return strings.Join([]string{"ExportCertificateRequest", string(data)}, " ")
}
