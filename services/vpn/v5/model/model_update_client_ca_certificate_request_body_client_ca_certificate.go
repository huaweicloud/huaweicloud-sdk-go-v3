package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateClientCaCertificateRequestBodyClientCaCertificate struct {

	// 客户端 CA 证书名
	Name *string `json:"name,omitempty"`
}

func (o UpdateClientCaCertificateRequestBodyClientCaCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateClientCaCertificateRequestBodyClientCaCertificate struct{}"
	}

	return strings.Join([]string{"UpdateClientCaCertificateRequestBodyClientCaCertificate", string(data)}, " ")
}
