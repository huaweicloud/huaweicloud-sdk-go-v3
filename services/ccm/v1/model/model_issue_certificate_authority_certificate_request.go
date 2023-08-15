package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueCertificateAuthorityCertificateRequest Request Object
type IssueCertificateAuthorityCertificateRequest struct {

	// 所要激活的从属CA证书ID。
	CaId string `json:"ca_id"`

	Body *IssueCertificateAuthorityCertificateRequestBody `json:"body,omitempty"`
}

func (o IssueCertificateAuthorityCertificateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCertificateAuthorityCertificateRequest struct{}"
	}

	return strings.Join([]string{"IssueCertificateAuthorityCertificateRequest", string(data)}, " ")
}
