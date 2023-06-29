package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableCertificateAuthorityRequest Request Object
type DisableCertificateAuthorityRequest struct {

	// 所要禁用CA证书ID。
	CaId string `json:"ca_id"`
}

func (o DisableCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"DisableCertificateAuthorityRequest", string(data)}, " ")
}
