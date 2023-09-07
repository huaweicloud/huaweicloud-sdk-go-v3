package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCertReq struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *CertificateKindObj `json:"kind"`

	Spec *UpdateSpecCert `json:"spec"`
}

func (o UpdateCertReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertReq struct{}"
	}

	return strings.Join([]string{"UpdateCertReq", string(data)}, " ")
}
