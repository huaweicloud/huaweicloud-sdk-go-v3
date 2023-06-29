package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestoreCertificateAuthorityRequest Request Object
type RestoreCertificateAuthorityRequest struct {

	// 所需要恢复的CA证书ID。
	CaId string `json:"ca_id"`
}

func (o RestoreCertificateAuthorityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreCertificateAuthorityRequest struct{}"
	}

	return strings.Join([]string{"RestoreCertificateAuthorityRequest", string(data)}, " ")
}
