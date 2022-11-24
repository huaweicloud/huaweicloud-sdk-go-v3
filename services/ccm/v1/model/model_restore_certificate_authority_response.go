package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type RestoreCertificateAuthorityResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o RestoreCertificateAuthorityResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreCertificateAuthorityResponse struct{}"
	}

	return strings.Join([]string{"RestoreCertificateAuthorityResponse", string(data)}, " ")
}
