package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainCertificateRequest Request Object
type ShowDomainCertificateRequest struct {

	// 域名id
	DomainId string `json:"domain_id"`
}

func (o ShowDomainCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainCertificateRequest struct{}"
	}

	return strings.Join([]string{"ShowDomainCertificateRequest", string(data)}, " ")
}
