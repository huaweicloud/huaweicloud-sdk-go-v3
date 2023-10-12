package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDetailsOfCertificateV2Request Request Object
type ShowDetailsOfCertificateV2Request struct {

	// 证书的编号
	CertificateId string `json:"certificate_id"`
}

func (o ShowDetailsOfCertificateV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfCertificateV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfCertificateV2Request", string(data)}, " ")
}
