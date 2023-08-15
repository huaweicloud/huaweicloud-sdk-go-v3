package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RevokeCertificateAuthorityResponse Response Object
type RevokeCertificateAuthorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RevokeCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RevokeCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"RevokeCertificateAuthorityResponse", string(data)}, " ")
}
