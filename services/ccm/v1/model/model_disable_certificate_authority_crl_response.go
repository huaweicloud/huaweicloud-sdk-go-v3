package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisableCertificateAuthorityCrlResponse Response Object
type DisableCertificateAuthorityCrlResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DisableCertificateAuthorityCrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisableCertificateAuthorityCrlResponse struct{}"
	}

	return strings.Join([]string{"DisableCertificateAuthorityCrlResponse", string(data)}, " ")
}
