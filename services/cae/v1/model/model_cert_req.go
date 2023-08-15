package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CertReq struct {

	// API版本。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Certificate”，该值不可修改。
	Kind string `json:"kind"`

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
