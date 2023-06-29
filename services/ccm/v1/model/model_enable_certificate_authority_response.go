package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCertificateAuthorityResponse Response Object
type EnableCertificateAuthorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"EnableCertificateAuthorityResponse", string(data)}, " ")
}
