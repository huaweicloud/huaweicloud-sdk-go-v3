package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCertificateAuthorityRequest Request Object
type ShowCertificateAuthorityRequest struct {

	// CA证书ID。
	CaId string `json:"ca_id"`
}

func (o ShowCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"ShowCertificateAuthorityRequest", string(data)}, " ")
}
