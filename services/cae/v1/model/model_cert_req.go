package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertReq struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *CertificateKindObj `json:"kind"`

	Metadata *CreateMetaCert `json:"metadata"`

	Spec *CreateSpecCert `json:"spec"`
}

func (o CertReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CertReq struct{}"
	}

	return strings.Join([]string{"CertReq", string(data)}, " ")
}
