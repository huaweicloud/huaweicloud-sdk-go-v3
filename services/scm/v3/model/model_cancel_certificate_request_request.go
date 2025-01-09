package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelCertificateRequestRequest Request Object
type CancelCertificateRequestRequest struct {

	// 证书ID。
	CertificateId string `json:"certificate_id"`
}

func (o CancelCertificateRequestRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelCertificateRequestRequest struct{}"
	}

	return strings.Join([]string{"CancelCertificateRequestRequest", string(data)}, " ")
}
