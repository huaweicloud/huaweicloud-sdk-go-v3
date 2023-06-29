package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDomainResponse Response Object
type ShowDomainResponse struct {

	//
	ProjectName *string `json:"project_name,omitempty"`

	//
	ProjectId *string `json:"project_id,omitempty"`

	//
	DomainId *string `json:"domain_id,omitempty"`

	//
	DomainName     *string `json:"domain_name,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowDomainResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDomainResponse struct{}"
	}

	return strings.Join([]string{"ShowDomainResponse", string(data)}, " ")
}
