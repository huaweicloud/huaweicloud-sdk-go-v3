package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCertificateAuthorityRequest Request Object
type EnableCertificateAuthorityRequest struct {

	// 所要启用的CA证书ID。
	CaId string `json:"ca_id"`
}

func (o EnableCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"EnableCertificateAuthorityRequest", string(data)}, " ")
}
