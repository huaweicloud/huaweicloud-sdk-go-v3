package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportCertCsrFileRequest Request Object
type ExportCertCsrFileRequest struct {
	Body *CreateCertSignatureReq `json:"body,omitempty"`
}

func (o ExportCertCsrFileRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportCertCsrFileRequest struct{}"
	}

	return strings.Join([]string{"ExportCertCsrFileRequest", string(data)}, " ")
}
