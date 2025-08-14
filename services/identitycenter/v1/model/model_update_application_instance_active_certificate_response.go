package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateApplicationInstanceActiveCertificateResponse Response Object
type UpdateApplicationInstanceActiveCertificateResponse struct {
	ApplicationInstanceCertificate *CertificateDto `json:"application_instance_certificate,omitempty"`
	HttpStatusCode                 int             `json:"-"`
}

func (o UpdateApplicationInstanceActiveCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateApplicationInstanceActiveCertificateResponse struct{}"
	}

	return strings.Join([]string{"UpdateApplicationInstanceActiveCertificateResponse", string(data)}, " ")
}
