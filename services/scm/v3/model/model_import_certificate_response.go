package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ImportCertificateResponse struct {
	// 证书id。

	CertificateId  *string `json:"certificate_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ImportCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificateResponse struct{}"
	}

	return strings.Join([]string{"ImportCertificateResponse", string(data)}, " ")
}
