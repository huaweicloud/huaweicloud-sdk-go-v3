package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCertificateAuthorityCrlRequest Request Object
type EnableCertificateAuthorityCrlRequest struct {

	// 所要启用CRL的CA证书ID。
	CaId string `json:"ca_id"`

	Body *EnableCertificateAuthorityCrlRequestBody `json:"body,omitempty"`
}

func (o EnableCertificateAuthorityCrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCertificateAuthorityCrlRequest struct{}"
	}

	return strings.Join([]string{"EnableCertificateAuthorityCrlRequest", string(data)}, " ")
}
