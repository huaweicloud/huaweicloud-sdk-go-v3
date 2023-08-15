package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableCertificateAuthorityResponse Response Object
type DisableCertificateAuthorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"DisableCertificateAuthorityResponse", string(data)}, " ")
}
