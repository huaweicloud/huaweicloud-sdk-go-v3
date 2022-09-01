package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateProjectDomainResponse struct {

	// 领域名称
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 领域id
	DomainId       *string `json:"domain_id,omitempty" xml:"domain_id"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateProjectDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateProjectDomainResponse struct{}"
	}

	return strings.Join([]string{"CreateProjectDomainResponse", string(data)}, " ")
}
