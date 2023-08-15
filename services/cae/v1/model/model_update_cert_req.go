package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateCertReq struct {

	// API版本.
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Certificate”，该值不可修改。
	Kind string `json:"kind"`

	Spec *UpdateSpecCert `json:"spec"`
}

func (o UpdateCertReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateCertReq struct{}"
	}

	return strings.Join([]string{"UpdateCertReq", string(data)}, " ")
}
