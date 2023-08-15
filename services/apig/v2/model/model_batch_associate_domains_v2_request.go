package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchAssociateDomainsV2Request Request Object
type BatchAssociateDomainsV2Request struct {

	// 证书的编号
	CertificateId string `json:"certificate_id"`

	Body *AttachOrDetachDomainsReqBody `json:"body,omitempty"`
}

func (o BatchAssociateDomainsV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchAssociateDomainsV2Request struct{}"
	}

	return strings.Join([]string{"BatchAssociateDomainsV2Request", string(data)}, " ")
}
