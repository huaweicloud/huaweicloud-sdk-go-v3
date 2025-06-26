package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateDomainNameRequestBody struct {

	// SCM服务的证书ID
	CertificateId string `json:"certificate_id"`
}

func (o UpdateDomainNameRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateDomainNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateDomainNameRequestBody", string(data)}, " ")
}
