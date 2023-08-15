package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateProjectDomainResponse Response Object
type UpdateProjectDomainResponse struct {

	// 领域名称
	DomainName *string `json:"domain_name,omitempty"`

	// 领域id
	DomainId       *string `json:"domain_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateProjectDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProjectDomainResponse struct{}"
	}

	return strings.Join([]string{"UpdateProjectDomainResponse", string(data)}, " ")
}
