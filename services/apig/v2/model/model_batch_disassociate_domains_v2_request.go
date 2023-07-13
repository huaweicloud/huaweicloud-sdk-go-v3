package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDisassociateDomainsV2Request Request Object
type BatchDisassociateDomainsV2Request struct {

	// 证书的编号
	CertificateId string `json:"certificate_id"`

	Body *AttachOrDetachDomainsReqBody `json:"body,omitempty"`
}

func (o BatchDisassociateDomainsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDisassociateDomainsV2Request struct{}"
	}

	return strings.Join([]string{"BatchDisassociateDomainsV2Request", string(data)}, " ")
}
