package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnsubscribeCertificateRequest Request Object
type UnsubscribeCertificateRequest struct {

	// 证书id。
	CertId string `json:"cert_id"`
}

func (o UnsubscribeCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnsubscribeCertificateRequest struct{}"
	}

	return strings.Join([]string{"UnsubscribeCertificateRequest", string(data)}, " ")
}
