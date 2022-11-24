package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type IssueCertificateAuthorityCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o IssueCertificateAuthorityCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueCertificateAuthorityCertificateResponse struct{}"
	}

	return strings.Join([]string{"IssueCertificateAuthorityCertificateResponse", string(data)}, " ")
}
