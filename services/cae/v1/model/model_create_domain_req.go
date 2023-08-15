package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDomainReq struct {

	// API版本，固定值“v1”，该值不可修改。
	ApiVersion string `json:"api_version"`

	// API类型，固定值“Domain”，该值不可修改。
	Kind string `json:"kind"`

	Metadata *CreateMetaDomain `json:"metadata"`
}

func (o CreateDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainReq struct{}"
	}

	return strings.Join([]string{"CreateDomainReq", string(data)}, " ")
}
