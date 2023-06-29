package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EnableCertificateAuthorityCrlResponse Response Object
type EnableCertificateAuthorityCrlResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o EnableCertificateAuthorityCrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnableCertificateAuthorityCrlResponse struct{}"
	}

	return strings.Join([]string{"EnableCertificateAuthorityCrlResponse", string(data)}, " ")
}
