package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
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
