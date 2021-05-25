package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type AssociateCertificateV2Request struct {
	// 实例编号

	InstanceId string `json:"instance_id"`
	// 域名的编号

	DomainId string `json:"domain_id"`
	// 分组的编号

	GroupId string `json:"group_id"`

	Body *DomainCertReq `json:"body,omitempty"`
}

func (o AssociateCertificateV2Request) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AssociateCertificateV2Request struct{}"
	}

	return strings.Join([]string{"AssociateCertificateV2Request", string(data)}, " ")
}
