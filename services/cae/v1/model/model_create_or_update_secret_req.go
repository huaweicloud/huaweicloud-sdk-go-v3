package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateOrUpdateSecretReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// 资源种类。
	Kind string `json:"kind"`

	Spec *CreateOrUpdateSecretDetail `json:"spec"`
}

func (o CreateOrUpdateSecretReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateOrUpdateSecretReq struct{}"
	}

	return strings.Join([]string{"CreateOrUpdateSecretReq", string(data)}, " ")
}
