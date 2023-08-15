package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCertificateAuthorityResponse Response Object
type CreateCertificateAuthorityResponse struct {

	// 当前签发的CA证书ID。
	CaId           *string `json:"ca_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"CreateCertificateAuthorityResponse", string(data)}, " ")
}
