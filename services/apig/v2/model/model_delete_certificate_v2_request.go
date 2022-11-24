package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteCertificateV2Request struct {

	// 证书的编号
	CertificateId string `json:"certificate_id"`
}

func (o DeleteCertificateV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateV2Request struct{}"
	}

	return strings.Join([]string{"DeleteCertificateV2Request", string(data)}, " ")
}
