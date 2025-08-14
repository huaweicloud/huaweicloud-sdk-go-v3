package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IdpCertificateBody struct {

	// 证书全局唯一标识符（ID）
	CertificateId string `json:"certificate_id"`

	// 证书状态
	Status string `json:"status"`
}

func (o IdpCertificateBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IdpCertificateBody struct{}"
	}

	return strings.Join([]string{"IdpCertificateBody", string(data)}, " ")
}
