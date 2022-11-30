package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowDetailsOfDomainNameCertificateV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 域名的编号
	DomainId string `json:"domain_id"`

	// 证书的编号
	CertificateId string `json:"certificate_id"`
}

func (o ShowDetailsOfDomainNameCertificateV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDetailsOfDomainNameCertificateV2Request struct{}"
	}

	return strings.Join([]string{"ShowDetailsOfDomainNameCertificateV2Request", string(data)}, " ")
}
