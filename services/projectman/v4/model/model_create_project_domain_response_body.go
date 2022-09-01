package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateProjectDomainResponseBody struct {

	// 领域名称
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 领域id
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`
}

func (o CreateProjectDomainResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectDomainResponseBody struct{}"
	}

	return strings.Join([]string{"CreateProjectDomainResponseBody", string(data)}, " ")
}
