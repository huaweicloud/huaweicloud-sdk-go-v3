package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClientCaCertificateRequestBody struct {
	ClientCaCertificate *UpdateClientCaCertificateRequestBodyClientCaCertificate `json:"client_ca_certificate,omitempty"`
}

func (o UpdateClientCaCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientCaCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateClientCaCertificateRequestBody", string(data)}, " ")
}
