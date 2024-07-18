package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportClientCaCertificateRequestBody struct {
	ClientCaCertificate *ImportClientCaCertificateRequestBodyClientCaCertificate `json:"client_ca_certificate,omitempty"`
}

func (o ImportClientCaCertificateRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportClientCaCertificateRequestBody struct{}"
	}

	return strings.Join([]string{"ImportClientCaCertificateRequestBody", string(data)}, " ")
}
