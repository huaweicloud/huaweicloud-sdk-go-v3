package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportCertificateAuthorityCertificateResponse Response Object
type ImportCertificateAuthorityCertificateResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o ImportCertificateAuthorityCertificateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportCertificateAuthorityCertificateResponse struct{}"
	}

	return strings.Join([]string{"ImportCertificateAuthorityCertificateResponse", string(data)}, " ")
}
