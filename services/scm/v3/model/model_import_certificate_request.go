package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCertificateRequest Request Object
type ImportCertificateRequest struct {
	Body *ImportCertificateRequestBody `json:"body,omitempty"`
}

func (o ImportCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificateRequest struct{}"
	}

	return strings.Join([]string{"ImportCertificateRequest", string(data)}, " ")
}
