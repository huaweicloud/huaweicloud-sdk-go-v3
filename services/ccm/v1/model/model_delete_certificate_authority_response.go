package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteCertificateAuthorityResponse Response Object
type DeleteCertificateAuthorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"DeleteCertificateAuthorityResponse", string(data)}, " ")
}
