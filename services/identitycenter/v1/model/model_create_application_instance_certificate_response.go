package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateApplicationInstanceCertificateResponse Response Object
type CreateApplicationInstanceCertificateResponse struct {
	ApplicationInstanceCertificate *CertificateDto `json:"application_instance_certificate,omitempty"`
	HttpStatusCode                 int             `json:"-"`
}

func (o CreateApplicationInstanceCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateApplicationInstanceCertificateResponse struct{}"
	}

	return strings.Join([]string{"CreateApplicationInstanceCertificateResponse", string(data)}, " ")
}
