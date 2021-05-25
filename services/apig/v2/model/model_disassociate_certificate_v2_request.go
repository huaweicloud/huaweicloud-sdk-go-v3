package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisassociateCertificateV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 域名的编号

	DomainId string `json:"domain_id"`
	// 分组的编号

	GroupId string `json:"group_id"`
	// 证书的编号

	CertificateId string `json:"certificate_id"`
}

func (o DisassociateCertificateV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisassociateCertificateV2Request struct{}"
	}

	return strings.Join([]string{"DisassociateCertificateV2Request", string(data)}, " ")
}
