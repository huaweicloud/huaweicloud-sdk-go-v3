package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportClientCaCertificateRequestBodyClientCaCertificate struct {

	// 证书名
	Name *string `json:"name,omitempty"`

	// 内容
	Content string `json:"content"`
}

func (o ImportClientCaCertificateRequestBodyClientCaCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportClientCaCertificateRequestBodyClientCaCertificate struct{}"
	}

	return strings.Join([]string{"ImportClientCaCertificateRequestBodyClientCaCertificate", string(data)}, " ")
}
