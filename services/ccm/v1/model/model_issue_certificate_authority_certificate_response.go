package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// IssueCertificateAuthorityCertificateResponse Response Object
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
