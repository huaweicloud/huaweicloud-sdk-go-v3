package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateDomainReq struct {
	ApiVersion *ApiVersionObj `json:"api_version"`

	Kind *DomainKindObj `json:"kind"`

	Metadata *CreateMetaDomain `json:"metadata"`
}

func (o CreateDomainReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDomainReq struct{}"
	}

	return strings.Join([]string{"CreateDomainReq", string(data)}, " ")
}
