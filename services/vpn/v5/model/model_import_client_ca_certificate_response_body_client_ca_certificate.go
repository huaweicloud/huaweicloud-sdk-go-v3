package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportClientCaCertificateResponseBodyClientCaCertificate 客户端CA证书
type ImportClientCaCertificateResponseBodyClientCaCertificate struct {

	// ID
	Id *string `json:"id,omitempty"`
}

func (o ImportClientCaCertificateResponseBodyClientCaCertificate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportClientCaCertificateResponseBodyClientCaCertificate struct{}"
	}

	return strings.Join([]string{"ImportClientCaCertificateResponseBodyClientCaCertificate", string(data)}, " ")
}
