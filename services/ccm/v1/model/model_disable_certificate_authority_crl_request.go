package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DisableCertificateAuthorityCrlRequest struct {

	// 所要禁用CRL的CA证书ID。
	CaId string `json:"ca_id"`
}

func (o DisableCertificateAuthorityCrlRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableCertificateAuthorityCrlRequest struct{}"
	}

	return strings.Join([]string{"DisableCertificateAuthorityCrlRequest", string(data)}, " ")
}
